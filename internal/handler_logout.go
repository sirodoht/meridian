package internal

import (
	"net/http"
	"time"
)

func (handlers *Handlers) HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Expires: time.Now().Add(-1 * time.Hour),
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
