package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	gsqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite"

	"github.com/sirodoht/meridian/internal"
)

func main() {
	debugMode := os.Getenv("DEBUG")

	databaseDSN := os.Getenv("DATABASE_DSN")
	if databaseDSN == "" {
		databaseDSN = ":memory:"
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // nolint: errcheck

	db, err := gorm.Open(
		gsqlite.Open(databaseDSN),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	store := internal.NewSQLStore(db)
	handlers := internal.NewHandlers(store, logger)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// middleware to check if user is authenticated
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var identityNRI string
			isAuthenticated := false
			c, err := r.Cookie("session")
			if err != nil {
				logger.Info("no session cookie")
			} else {
				session, err := store.GetSession(r.Context(), c.Value)
				if err != nil {
					logger.Info("failed to get session", zap.Error(err))
				} else {
					// TODO: check if session is expired
					identityNRI = session.IdentityNRI
					isAuthenticated = true
				}
			}
			ctx := context.WithValue(r.Context(), internal.KeyIdentity, identityNRI)
			ctx = context.WithValue(ctx, internal.KeyIsAuthenticated, isAuthenticated)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	// routes
	r.Get("/", handlers.RenderIndex)

	// static files
	if debugMode == "1" {
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
