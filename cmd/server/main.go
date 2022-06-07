package main

import (
	"context"
	"flag"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/tmrrwnxtsn/const-payments-api/internal/config"
	"github.com/tmrrwnxtsn/const-payments-api/internal/handler"
	"github.com/tmrrwnxtsn/const-payments-api/internal/server"
	"github.com/tmrrwnxtsn/const-payments-api/internal/service"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store/sqlstore"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var flagConfig = flag.String("config", "./configs/local.yml", "path to config file")

// @title           Constanta Payments API
// @version         1.0
// @description     Эмулятор платёжного сервиса, позволяющего работать с платежами (транзакциями).
// @termsOfService  http://swagger.io/terms/

// @license.name  The MIT License (MIT)
// @license.url   https://mit-license.org/

// @host      localhost:8080
// @BasePath  /api

func main() {
	flag.Parse()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := config.Load(*flagConfig)
	if err != nil {
		logger.Fatalf("failed to load config data: %s", err)
	}

	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logger.Fatalf("failed to set logging level: %s", err)
	}
	logger.SetLevel(level)

	db, err := sqlx.Connect("pgx", cfg.DSN)
	if err != nil {
		logger.Fatalf("failed to establish database connection %s", err)
	}

	st := sqlstore.NewStore(db)
	services := service.NewServices(st)
	router := handler.NewHandler(services, logger)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	srv := server.NewServer(cfg.BindAddr, router.InitRoutes())
	go func() {
		if err = srv.Run(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("error occurred while running server: %s", err)
		}
	}()
	logger.Infof("server is running at %v", cfg.BindAddr)

	<-quit
	logger.Info("server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		if err = db.Close(); err != nil {
			logger.Fatalf("failed to close the database connection: %s", err)
		}
		cancel()
	}()

	if err = srv.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown failed: %s", err)
	}

	logger.Info("server exited properly")
}
