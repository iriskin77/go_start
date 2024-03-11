package main

import (
	"log"

	gostart "github.com/iriskin77/go_start"
	"github.com/iriskin77/go_start/pkg/handler"
	"github.com/iriskin77/go_start/pkg/repository"
	"github.com/iriskin77/go_start/pkg/service"
)

func main() {
	// инициализируем экземпляр сервера с помощью ключевого слова new(). Благодаря New мы передаем указатель на сервер

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(gostart.Server)

	if err := server.Run("9090", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while ruuning http server %s", err.Error())
	}
}
