package main

import (
	"github.com/qsmsoft/blog/config"
	"github.com/qsmsoft/blog/internal/app"
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

	e := app.SetupApp(cfg)

	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
