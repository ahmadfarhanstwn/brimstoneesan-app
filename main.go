package main

import (
	"myapp/data"
	"myapp/handlers"

	"github.com/ahmadfarhanstwn/brimstoneesan"
)

type application struct {
	App      *brimstoneesan.Brimstoneesan
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	b := initApplication()
	b.App.ListenAndServer()
}
