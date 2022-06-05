package main

import (
	"flag"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/tmrrwnxtsn/const-payments-api/internal/config"
	"github.com/tmrrwnxtsn/const-payments-api/internal/service"
	"github.com/tmrrwnxtsn/const-payments-api/internal/store"
	logging "github.com/tmrrwnxtsn/const-payments-api/pkg/log"
	"log"
	"os"
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
	defer db.Close()

	st := store.NewStore(db, logger)
	_ = service.NewService(st, logger)
}
