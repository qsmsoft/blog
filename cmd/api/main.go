package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/qsmsoft/blog/config"
	"github.com/qsmsoft/blog/internal/server"
	"github.com/qsmsoft/blog/pkg/db/postgres"
	"github.com/qsmsoft/blog/pkg/logger"
	"github.com/qsmsoft/blog/pkg/utils"
	"log"
	"os"
)

func main() {

	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s", cfg.Server.AppVersion, cfg.Logger.Level)

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer func(psqlDB *sqlx.DB) {
		err := psqlDB.Close()
		if err != nil {
			appLogger.Fatalf("Postgresql close: %s", err)
		}
	}(psqlDB)

	s := server.NewServer(cfg, psqlDB, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}

}
