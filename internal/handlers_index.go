package internal

import (
	"html/template"
	"net/http"
)

func (handlers *Handlers) HandleIndex(
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
