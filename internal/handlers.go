package internal

import (
	"html/template"
	"net/http"

	"go.uber.org/zap"
)

type Handlers struct {
	logger *zap.Logger
	api    API
}

func NewHandlers(
	logger *zap.Logger,
	api API,
) *Handlers {
	return &Handlers{
		logger: logger,
		api:    api,
	}
}

func (handlers *Handlers) RenderIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(
		"internal/templates/layout.html",
		"internal/templates/index.html",
	)
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, map[string]interface{}{
		"IsAuthenticated": r.Context().Value(KeyIsAuthenticated),
		"Username":        r.Context().Value(KeyIdentity),
	})
	if err != nil {
		panic(err)
	}
	handlers.logger.Info("render index success")
}
