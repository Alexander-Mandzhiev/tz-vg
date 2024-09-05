package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"tz-vg/internal/apiserver"
	"tz-vg/internal/config"
	"tz-vg/internal/handlers"
	"tz-vg/internal/repository"
	"tz-vg/internal/service"
	"tz-vg/pkg/logger"
	"tz-vg/pkg/postgres"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger()
	log.Info("Starting app", slog.String("port", cfg.Address))
	log.Debug("Debug messages are enabled")

	db, err := postgres.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Error("failed to initialize db: %s", logger.Err(err))
	}
	log.Debug("Initialize database")

	repository := repository.New(db)
	log.Info("Initialize repository")
	service := service.NewService(repository)
	log.Info("Initialize service")
	ctx := context.Background()

	handler := handlers.NewHandler(service, log, ctx)
	log.Info("initialize handlers")

	srv := new(apiserver.APIServer)

	go func() {
		if err := srv.Start(cfg, handler.InitRouters()); err != nil {
			log.Error("error occured while running http server: %s", logger.Err(err))
		}
	}()
	log.Info("starting application")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("stopping application", slog.String("signal", sign.String()))
	srv.Shutdown(ctx)
	log.Info("application stopped")
	repository.Close()
	log.Info("db connection close")
}
