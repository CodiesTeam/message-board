package user

import (
	"fmt"
	"message-board/server/mbmux"
	"net/http"
	"strconv"
)

// Serve startup all serve handlers for handling user.
func Serve() {
	mbmux.SharedMux.HandleFunc("/register", register)
}

func register(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Expected code %d, got %d: %s", http.StatusAccepted, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	var name, email, password string
	var gender Gender
	name, email, password = r.Form.Get("name"), r.Form.Get("email"), r.Form.Get("password")
	genderInt, e := strconv.Atoi(r.Form.Get("gender"))
	gender = Gender(genderInt)
	if e != nil || (gender != Man && gender != Woman) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Expected code %d, got %d: %s", http.StatusAccepted, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Expected code %d, got %d: %s", http.StatusAccepted, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Expected code %d, got %d: %s", http.StatusAccepted, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Expected code %d, got %d: %s", http.StatusAccepted, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	_, e = CreateUser(name, email, password, gender)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Expected code %d, got %d: %s", http.StatusAccepted, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}
