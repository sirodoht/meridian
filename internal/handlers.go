package internal

import (
	"context"
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

//go:embed static
var staticFiles embed.FS

//go:embed templates
var templateFiles embed.FS

type Handlers struct {
	logger *zap.Logger
	api    API
	store  Store
}

type TemplateValues struct {
	User  *User
	Error string
}

func NewHandlers(
	logger *zap.Logger,
	api API,
	store Store,
) *Handlers {
	return &Handlers{
		logger: logger,
		api:    api,
		store:  store,
	}
}

func (handlers *Handlers) Register(r *chi.Mux) {
	// register middleware
	r.Use(handlers.authMiddleware)

	// register handlers
	r.Get("/", handlers.HandleIndex)
	r.Get("/login", handlers.HandleLogin)
	r.Post("/login", handlers.HandleLogin)
	r.Post("/logout", handlers.HandleLogout)
	r.Get("/signup", handlers.HandleRegister)
	r.Post("/signup", handlers.HandleRegister)
	r.Get("/notes", handlers.HandleNotes)
	r.Get("/notes/new", handlers.HandleNotesNew)
	r.Post("/notes/new", handlers.HandleNotesNew)
	r.Post("/follow", handlers.HandleFollow)
	r.Get("/profile/{keygraphID}", handlers.HandleProfile)

	// static files
	fs := http.FileServer(http.FS(staticFiles))
	r.Handle("/static/*", fs)
}

// middleware to check if user is authenticated
func (handlers *Handlers) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string
		isAuthenticated := false
		c, err := r.Cookie("session")
		if err != nil {
			handlers.logger.Info("no session cookie")
		} else {
			session, err := handlers.store.GetSession(r.Context(), c.Value)
			if err != nil {
				handlers.logger.Info("failed to get session", zap.Error(err))
			} else {
				// TODO: check if session is expired
				username = session.Username
				isAuthenticated = true
			}
		}
		ctx := context.WithValue(r.Context(), KeyUsername, username)
		ctx = context.WithValue(ctx, KeyIsAuthenticated, isAuthenticated)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (handlers *Handlers) valuesFromCtx(ctx context.Context) TemplateValues {
	values := TemplateValues{}
	username, ok := ctx.Value(KeyUsername).(string)
	if ok {
		user, err := handlers.store.GetUser(ctx, username)
		if err != nil {
			// TODO: handle error
		}
		values.User = user
	}

	return values
}
