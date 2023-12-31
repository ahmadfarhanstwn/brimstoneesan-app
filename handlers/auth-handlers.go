package handlers

import (
	"net/http"
)

func (h *Handlers) UserLogin(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "login", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) PostUserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := h.Models.Users.GetByEmail(email)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	isMatched, err := user.PasswordMatched(password)
	if err != nil {
		w.Write([]byte("Error validating password"))
		return
	}

	if !isMatched {
		w.Write([]byte("Password is incorrect"))
		return
	}

	h.App.Session.Put(r.Context(), "userId", user.ID)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
