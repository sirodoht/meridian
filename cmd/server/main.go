package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moul.io/chizap"

	"nimona.io"

	"github.com/sirodoht/meridian/internal"
)

func main() {
	debugMode, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	databaseDSN := os.Getenv("DATABASE_DSN")
	if databaseDSN == "" {
		databaseDSN = "meridian.sqlite"
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // nolint: errcheck

	db, err := gorm.Open(
		sqlite.Open(databaseDSN),
		&gorm.Config{},
	)
	if err != nil {
		logger.Fatal("failed to open database", zap.Error(err))
	}

	// Enable debug mode
	if debugMode {
		db = db.Debug()
	}

	// Construct a new document store
	docStore, err := nimona.NewDocumentStore(db)
	if err != nil {
		logger.Fatal("failed to create document store", zap.Error(err))
	}

	// Construct a new identity store
	idStore, err := nimona.NewIdentityStore(db)
	if err != nil {
		logger.Fatal("failed to create identity store", zap.Error(err))
	}

	// Construct a new meridian store
	meridianStore := internal.NewSQLStore(db)

	// Construct a new meridian api
	api := internal.NewAPI(logger, meridianStore, docStore, idStore)

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
	fmt.Println("Listening on http://127.0.0.1:8000/")
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
