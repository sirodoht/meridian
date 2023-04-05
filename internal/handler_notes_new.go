package internal

import (
	"html/template"
	"net/http"
)

func (handlers *Handlers) HandleNotesNew(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(
		"internal/templates/layout.html",
		"internal/templates/notes_new.html",
	)
	if err != nil {
		panic(err)
	}

	values := handlers.valuesFromCtx(r.Context())

	if r.Method == http.MethodPost {
		r.ParseForm()
		body := r.Form.Get("body")
		_, err := handlers.api.CreateNote(
			r.Context(),
			&CreateNoteRequest{
				Username: values.User.Username,
				Content:  body,
			},
		)
		if err == nil {
			http.Redirect(w, r, "/notes", http.StatusSeeOther)
			return
		}
		values.Error = err.Error()
	}

	err = t.Execute(w, values)
	if err != nil {
		panic(err)
	}
}
