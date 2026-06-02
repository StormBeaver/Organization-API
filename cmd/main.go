package main

import (
	"context"
	"orgService/internal/config"
	"orgService/internal/database"
	handler "orgService/internal/handlers"
	appLogger "orgService/internal/logger"
	"orgService/internal/repo"
	"orgService/internal/server"
	"orgService/internal/service"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := config.ReadConfig("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}

	cfg := config.GetConfigInstance()

	logger := appLogger.LogInit(cfg.Project.Debug)

	db, err := database.ConfigureGorm(&cfg, &logger)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init PosrgreSQL")
	}

	repo := repo.NewRepo(db)
	service := service.NewService(repo, &logger)
	handler := handler.NewHandler(service, &logger)
	server := server.NewServer(cfg.Rest.Host, cfg.Rest.Port, handler.Handler(), &logger)

	go server.Run()
	server.Shutdown(ctx)
}
