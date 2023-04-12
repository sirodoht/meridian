package internal

import (
	"html/template"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5"
)

func (handlers *Handlers) HandleProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(
		"internal/templates/layout.html",
		"internal/templates/profile.html",
	)
	if err != nil {
		panic(err)
	}

	// TODO(geoah): normalize all identities to not use NRIs
	identity := "nimona://id:" + chi.URLParam(r, "identity")

	followers, err := handlers.api.GetFollowers(r.Context(), &GetFollowersRequest{
		IdentityNRI: identity,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followees, err := handlers.api.GetFollowees(r.Context(), &GetFolloweesRequest{
		IdentityNRI: identity,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	spew.Dump(followers)
	spew.Dump(followees)

	values := struct {
		TemplateValues
		Followers []string
		Followees []string
	}{
		TemplateValues: handlers.valuesFromCtx(r.Context()),
		Followers:      followers.Followers,
		Followees:      followees.Followees,
	}

	err = t.Execute(w, values)
	if err != nil {
		panic(err)
	}
}
