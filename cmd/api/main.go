package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/qsmsoft/blog/config"
	"github.com/qsmsoft/blog/internal/server"
	"github.com/qsmsoft/blog/pkg/db/aws"
	"github.com/qsmsoft/blog/pkg/db/postgres"
	rClient "github.com/qsmsoft/blog/pkg/db/redis"
	"github.com/qsmsoft/blog/pkg/logger"
	"github.com/qsmsoft/blog/pkg/utils"
	"github.com/redis/go-redis/v9"
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

	redisClient := rClient.NewRedisClient(cfg)
	defer func(redisClient *redis.Client) {
		err := redisClient.Close()
		if err != nil {
			appLogger.Fatalf("Redis close: %s", err)
		}
	}(redisClient)
	appLogger.Info("Redis connected")

	awsClient, err := aws.NewAWSClient(cfg.AWS.Endpoint, cfg.AWS.MinioAccessKey, cfg.AWS.MinioSecretKey, cfg.AWS.UseSSL)
	if err != nil {
		appLogger.Errorf("AWS Client init: %s", err)
	}
	appLogger.Info("AWS S3 connected")

	s := server.NewServer(cfg, psqlDB, redisClient, awsClient, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}

}
