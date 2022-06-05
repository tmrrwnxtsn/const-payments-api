package main

import (
	"context"
	"flag"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/tmrrwnxtsn/const-payments-api/internal/config"
	"github.com/tmrrwnxtsn/const-payments-api/internal/handler"
	"github.com/tmrrwnxtsn/const-payments-api/internal/server"
	"github.com/tmrrwnxtsn/const-payments-api/internal/service"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
	logging "github.com/tmrrwnxtsn/const-payments-api/pkg/log"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var flagConfig = flag.String("config", "./configs/local.yaml", "path to config file")

func main() {
	flag.Parse()

	cfg, err := config.Load(*flagConfig)
	if err != nil {
		log.Fatalf("failed to load config data: %s", err)
	}

	logger := logging.New()
	if err = logger.SetLoggingLevel(cfg.LogLevel); err != nil {
		logger.Errorf("failed to set logging level: %s", err)
		os.Exit(-1)
	}

	db, err := sqlx.Connect("pgx", cfg.DSN)
	if err != nil {
		logger.Errorf("failed to establish database connection %s", err)
		os.Exit(-1)
	}

	st := store.NewStore(db, logger)
	serv := service.NewService(st, logger)
	router := handler.NewHandler(serv, logger)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	srv := server.NewServer(cfg.BindAddr, router.InitRoutes())
	go func() {
		if err = srv.Run(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("error occurred while running server: %s", err)
			os.Exit(-1)
		}
	}()
	logger.Infof("server is running at %v", cfg.BindAddr)

	<-quit
	logger.Info("server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		if err = db.Close(); err != nil {
			logger.Errorf("failed to close the database connection: %s", err)
			os.Exit(-1)
		}
		cancel()
	}()

	if err = srv.Shutdown(ctx); err != nil {
		logger.Errorf("server shutdown failed: %s", err)
		os.Exit(-1)
	}

	logger.Info("server exited properly")
}
