package internal

import (
	"context"
	"html/template"
	"net/http"

	"go.uber.org/zap"
)

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

func (handlers *Handlers) RenderIndex(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(
		"internal/templates/layout.html",
		"internal/templates/index.html",
	)
	if err != nil {
		panic(err)
	}

	values := handlers.valuesFromCtx(r.Context())
	err = t.Execute(w, values)
	if err != nil {
		panic(err)
	}
	handlers.logger.Info("render index success")
}
