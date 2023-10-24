package main

import (
	"log"
	"myapp/handlers"
	"os"

	"github.com/ahmadfarhanstwn/brimstoneesan"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//init brimstoneesan
	brims := &brimstoneesan.Brimstoneesan{}
	err = brims.New(path)
	if err != nil {
		log.Fatal(err)
	}

	brims.AppName = "myApp"
	brims.Debug = true

	myHandlers := &handlers.Handlers{
		App: brims,
	}

	app := &application{
		App:      brims,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
