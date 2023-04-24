package internal

import (
	"net/http"

	"go.uber.org/zap"

	"nimona.io"
)

func (handlers *Handlers) HandleFollow(w http.ResponseWriter, r *http.Request) {
	values := handlers.valuesFromCtx(r.Context())

	if values.User.KeygraphID.IsEmpty() {
		handlers.logger.Error("user not logged in, or no identity in context")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	r.ParseForm()
	followee := r.Form.Get("followee")
	followeeID, err := nimona.ParseKeygraphID(followee)
	_, err = handlers.api.Follow(r.Context(), &FollowRequest{
		KeygraphID:       values.User.KeygraphID,
		FolloweeIdentity: followeeID,
	})
	if err != nil {
		handlers.logger.Error("error following user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: redirect to the user's profile page or something
	http.Redirect(w, r, "/", http.StatusFound)
}
