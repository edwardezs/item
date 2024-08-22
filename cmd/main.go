package main

import (
	"context"
	"test/internal/config"
	"test/internal/handlers"
	"test/internal/repo"
	"test/internal/server"
	"test/internal/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const appName = "test"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := config.Get()
	if err != nil {
		logrus.Fatalf("failed to get config: %s", err.Error())
	}

	db, err := repo.NewPostgresDB(repo.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.User,
		Password: cfg.Pass,
		DBName:   cfg.Name,
		SSLMode:  cfg.SSLMode,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repo.NewRepository(db)
	service := service.NewService(repos)
	handlers := handlers.NewHandler(service)

	srv := server.New(cfg.ServerPort, handlers.InitRoutes())
	if err := srv.Run(ctx); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(appName+" gracefully finished")
}
