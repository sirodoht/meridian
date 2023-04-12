package internal

import (
	"html/template"
	"net/http"
)

func (handlers *Handlers) HandleRegister(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(
		"internal/templates/layout.html",
		"internal/templates/register.html",
	)
	if err != nil {
		panic(err)
	}

	values := TemplateValues{}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		email := r.Form.Get("email")
		res, err := handlers.api.Register(
			r.Context(),
			&RegisterRequest{
				Username: username,
				Password: password,
				Email:    email,
			},
		)
		if err == nil {
			http.SetCookie(w, &http.Cookie{
				Name:  "session",
				Value: res.SessionID,
			})
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		values.Error = err.Error()
	}

	err = t.Execute(w, values)
	if err != nil {
		panic(err)
	}
}
