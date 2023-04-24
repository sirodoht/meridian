package internal

import (
	"html/template"
	"net/http"
	"strings"
)

func (handlers *Handlers) HandleNotes(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.New("layout.html").
		Funcs(template.FuncMap{
			"identityFromNRI": func(s string) string {
				return strings.TrimPrefix(s, "nimona://id:")
			},
		}).ParseFiles(
		"internal/templates/layout.html",
		"internal/templates/notes.html",
	)
	if err != nil {
		panic(err)
	}

	req := &GetNotesRequest{}
	res, err := handlers.api.GetNotes(r.Context(), req)
	if err != nil {
		panic(err)
	}

	values := struct {
		TemplateValues
		Notes []*Note
	}{
		TemplateValues: handlers.valuesFromCtx(r.Context()),
		Notes:          res.Notes,
	}

	err = t.Execute(w, values)
	if err != nil {
		panic(err)
	}
}
