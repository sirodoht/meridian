package main

import (
	"context"
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
	r.Use(middleware.Logger)

	// middleware to check if user is authenticated
	// TODO: move to internal/middleware.go
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var username string
			isAuthenticated := false
			c, err := r.Cookie("session")
			if err != nil {
				logger.Info("no session cookie")
			} else {
				session, err := meridianStore.GetSession(r.Context(), c.Value)
				if err != nil {
					logger.Info("failed to get session", zap.Error(err))
				} else {
					// TODO: check if session is expired
					username = session.Username
					isAuthenticated = true
				}
			}
			ctx := context.WithValue(r.Context(), internal.KeyUsername, username)
			ctx = context.WithValue(ctx, internal.KeyIsAuthenticated, isAuthenticated)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	// routes
	// TODO: move to internal/router.go
	r.Get("/", handlers.RenderIndex)
	r.Get("/login", handlers.HandleLogin)
	r.Post("/login", handlers.HandleLogin)
	r.Get("/signup", handlers.HandleRegister)
	r.Post("/signup", handlers.HandleRegister)
	r.Get("/notes/new", handlers.HandleNotesNew)
	r.Post("/notes/new", handlers.HandleNotesNew)

	// static files
	if debugMode {
		fileServer := http.FileServer(http.Dir("./static/"))
		r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	}

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
