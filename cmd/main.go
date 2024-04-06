package main

import (
	"microservice"
	"microservice/config"
	"microservice/pkg/handlers"
	"microservice/pkg/repository"
	"microservice/pkg/service"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := config.Get()
	if err != nil {
		logrus.Fatalf("failed to get config: %s", err.Error())
		os.Exit(1)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.User,
		Password: cfg.Pass,
		DBName:   cfg.Name,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handlers.NewHandler(service)

	srv := new(microservice.Server)
	if err := srv.Run(cfg.ServerPort, handlers.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}
