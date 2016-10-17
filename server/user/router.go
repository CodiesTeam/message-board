package user

import (
	"fmt"
	"message-board/server/mbmux"
	"net/http"
)

// Serve startup all serve handlers for handling user.
func Serve() {
	mbmux.SharedMux.Handle("/register", register)
}

func register(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Expected code %d, got %d: %s", http.StatusAccepted, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	for k, v := range r.Form {

	}
}
