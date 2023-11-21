package main

import (
	"context"
	"log"

	"todo-api/internal/todo"
)

func main() {
	app := todo.NewApp()
	if err := app.Init(context.Background()); err != nil {
		log.Fatalf("%+v\n", err)
	}

	if err := app.Start(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}
