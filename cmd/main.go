package main

import (
	"log"

	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf(" problems with runing: %s", err.Error())
	}
}
