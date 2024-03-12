package main

import (
	"log"

	gostart "github.com/iriskin77/go_start"
	"github.com/iriskin77/go_start/pkg/handler"
	"github.com/iriskin77/go_start/pkg/repository"
	"github.com/iriskin77/go_start/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initialization config %s", err.Error())
	}

	// Инициализируем БД
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatal("failed to initialize db: %s", err.Error())
	}

	// инициализируем экземпляр сервера с помощью ключевого слова new(). Благодаря New мы передаем указатель на сервер

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(gostart.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while ruuning http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("/home/abc/Рабочий стол/go_start/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
