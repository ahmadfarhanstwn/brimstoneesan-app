package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	//middleware

	//routes
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	a.App.Routes.Get("/sessions-page", a.Handlers.TestSessionsPage)

	a.App.Routes.Get("/test-jet", func(w http.ResponseWriter, r *http.Request) {
		a.App.Render.JetPage(w, r, "test-jet", nil, nil)
	})

	a.App.Routes.Get("/test-database", func(w http.ResponseWriter, r *http.Request) {
		query := "select id, name from users where id = 1"
		row := a.App.Database.Pool.QueryRowContext(r.Context(), query)

		var id int
		var name string
		err := row.Scan(&id, &name)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%d %s", id, name)
	})

	//static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
