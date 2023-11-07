package main

import (
	"fmt"
	"myapp/data"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	//middleware

	//routes
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	a.App.Routes.Get("/sessions-page", a.Handlers.TestSessionsPage)

	a.App.Routes.Get("/test-insert", func(w http.ResponseWriter, r *http.Request) {
		user := data.User{
			FirstName: "Joko",
			LastName:  "Netanyahu",
			Email:     "netanyahu@bangsatkau.com",
			Active:    1,
			Password:  "netanyahubangsat",
		}

		id, err := a.Models.Users.Insert(user)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%d : %s", id, user.FirstName)
	})

	a.App.Routes.Get("/test-get-all", func(w http.ResponseWriter, r *http.Request) {
		users, err := a.Models.Users.GetAll()
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		for _, x := range users {
			fmt.Fprintf(w, x.LastName)
		}
	})

	a.App.Routes.Get("/test-get/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		user, err := a.Models.Users.GetById(id)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%s %s %s", user.FirstName, user.LastName, user.Email)
	})

	a.App.Routes.Get("/test-update/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		user, err := a.Models.Users.GetById(id)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		user.LastName = a.App.GenerateRandomString(10)
		err = user.Update(*user)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "updated last name to %s", user.LastName)
	})

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
