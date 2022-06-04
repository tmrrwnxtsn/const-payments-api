package main

import (
	"flag"
	"fmt"
	"github.com/tmrrwnxtsn/const-payments-api/internal/config"
	"github.com/tmrrwnxtsn/const-payments-api/pkg/log"
	"os"
)

var flagConfig = flag.String("config", "./configs/local.yaml", "path to config file")

func main() {
	flag.Parse()

	logger := log.New()

	cfg, err := config.Load(*flagConfig)
	if err != nil {
		logger.Errorf("failed to load config data: %s", err)
		os.Exit(-1)
	}

	fmt.Println(cfg.DSN, cfg.BindAddr)
}
