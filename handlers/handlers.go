package handlers

import (
	"myapp/data"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/ahmadfarhanstwn/brimstoneesan"
)

type Handlers struct {
	App    *brimstoneesan.Brimstoneesan
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) GoPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(w, r, "home", nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) JetPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(w, r, "jet-template", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) TestSessionsPage(w http.ResponseWriter, r *http.Request) {
	h.App.Session.Put(r.Context(), "foo", "bar")
	fooSession := h.App.Session.GetString(r.Context(), "foo")

	vars := make(jet.VarMap)
	vars.Set("foo", fooSession)
	err := h.App.Render.JetPage(w, r, "test-sessions", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
