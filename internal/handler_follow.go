package internal

import (
	"net/http"

	"go.uber.org/zap"
)

func (handlers *Handlers) HandleFollow(w http.ResponseWriter, r *http.Request) {
	values := handlers.valuesFromCtx(r.Context())

	if values.User.IdentityNRI == "" {
		handlers.logger.Error("user not logged in, or no identity in context")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	r.ParseForm()
	followee := r.Form.Get("followee")
	_, err := handlers.api.Follow(r.Context(), &FollowRequest{
		IdentityNRI:         values.User.IdentityNRI,
		FolloweeIdentityNRI: followee,
	})
	if err != nil {
		handlers.logger.Error("error following user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: redirect to the user's profile page or something
	http.Redirect(w, r, "/", http.StatusFound)
}
