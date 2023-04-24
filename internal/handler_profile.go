package internal

import (
	"html/template"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5"

	"nimona.io"
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

	keygraphID, err := nimona.ParseKeygraphID(chi.URLParam(r, "keygraphID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile, err := handlers.api.GetProfile(
		r.Context(),
		&GetProfileRequest{
			KeygraphID: keygraphID,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followers, err := handlers.api.GetFollowers(
		r.Context(),
		&GetFollowersRequest{
			KeygraphID: keygraphID,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followees, err := handlers.api.GetFollowees(
		r.Context(),
		&GetFolloweesRequest{
			KeygraphID: keygraphID,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	spew.Dump(profile.Profile)
	spew.Dump(followers)
	spew.Dump(followees)

	values := struct {
		TemplateValues
		Profile   Profile
		Followers []nimona.KeygraphID
		Followees []nimona.KeygraphID
	}{
		TemplateValues: handlers.valuesFromCtx(r.Context()),
		Profile:        profile.Profile,
		Followers:      followers.Followers,
		Followees:      followees.Followees,
	}

	err = t.Execute(w, values)
	if err != nil {
		panic(err)
	}
}
