package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/chizap"
	"moul.io/zapgorm2"

	"nimona.io"

	"github.com/sirodoht/meridian/internal"
)

type Config struct {
	Debug       bool   `envconfig:"DEBUG" default:"false"`
	Environment string `envconfig:"ENV" default:"development"`
	BindAddress string `envconfig:"BIND_ADDRESS" default:":8000"`
}

func main() {
	// Parse environment variables
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Normalize environment
	cfg.Environment = strings.ToLower(cfg.Environment)

	databaseDSN := os.Getenv("DATABASE_DSN")
	if databaseDSN == "" {
		databaseDSN = "meridian.sqlite"
	}

	// Construct a new zap logger
	logLevel := zap.NewAtomicLevel()
	logConfig := zap.NewProductionConfig()
	logEncoder := zapcore.NewJSONEncoder(logConfig.EncoderConfig)
	if cfg.Environment == "development" {
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logEncoder = zapcore.NewConsoleEncoder(logConfig.EncoderConfig)
	}
	logger := zap.New(
		zapcore.NewCore(
			logEncoder,
			zapcore.Lock(os.Stdout),
			logLevel,
		),
	)
	defer logger.Sync() // nolint: errcheck

	// Construct a new gorm logger
	zaplogger := zapgorm2.New(
		logger.Named("gorm"),
	)
	zaplogger.SetAsDefault()
	zaplogger.LogLevel = gormlogger.Info

	// Construct a new gorm database
	db, err := gorm.Open(
		sqlite.Open(databaseDSN),
		&gorm.Config{
			Logger: zaplogger,
		},
	)
	if err != nil {
		logger.Fatal("failed to open database", zap.Error(err))
	}

	// Enable debug mode
	if cfg.Debug {
		logger.Info("debug mode enabled")
		db = db.Debug()
		logLevel.SetLevel(zap.DebugLevel)
	}

	// Construct a new document store
	docStore, err := nimona.NewDocumentStore(db)
	if err != nil {
		logger.Fatal("failed to create document store", zap.Error(err))
	}

	// Construct a new keygraph store
	kgStore, err := nimona.NewKeygraphStore(db)
	if err != nil {
		logger.Fatal("failed to create identity store", zap.Error(err))
	}

	// Construct a new meridian store
	meridianStore := internal.NewSQLStore(db)

	// Construct a new meridian api
	api := internal.NewAPI(logger, meridianStore, docStore, kgStore)

	// Construct a new meridian router
	handlers := internal.NewHandlers(logger, api, meridianStore)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(chizap.New(logger, &chizap.Opts{
		WithReferer:   true,
		WithUserAgent: true,
	}))

	// register handlers
	handlers.Register(r)

	// serve
	logger.
		With(zap.String("address", cfg.BindAddress)).
		Info("starting server")

	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.BindAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
