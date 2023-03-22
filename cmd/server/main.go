package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sirodoht/meridian/internal"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func main() {
	debugMode := os.Getenv("DEBUG")

	databaseDSN := os.Getenv("DATABASE_DSN")
	db, err := sqlx.Open("sqlite", databaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // nolint: errcheck

	store := internal.NewSQLStore(db)
	handlers := internal.NewHandlers(store, logger)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// middleware to check if user is authenticated
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var username string
			isAuthenticated := false
			c, err := r.Cookie("session")
			if err != nil {
				fmt.Println(err)
			} else {
				username = store.GetUsernameSession(r.Context(), c.Value)
				if err == nil {
					isAuthenticated = true
				}
			}
			ctx := context.WithValue(r.Context(), internal.KeyUsername, username)
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
