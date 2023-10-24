package main

import (
	"myapp/handlers"

	"github.com/ahmadfarhanstwn/brimstoneesan"
)

type application struct {
	App      *brimstoneesan.Brimstoneesan
	Handlers *handlers.Handlers
}

func main() {
	b := initApplication()
	b.App.ListenAndServer()
}
