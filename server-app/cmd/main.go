package main

import (
	"mtg/internal/apiserver"
	"mtg/internal/config"
	"mtg/pkg/logging"

	"github.com/joho/godotenv"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("запуск приложения")

	logger.Debug("парсим конфиг")
	if err := godotenv.Load(); err != nil {
		logger.Info("файл .env не найден")
	}

	cfg := config.GetConfig()
	logger.Infof("конфиг получен: port: %s, DB_url: %s",
		cfg.Port, cfg.DBURL)

	logger.Debug("запускаем сервер")
	if err := apiserver.Start(cfg, logger); err != nil {
		logger.Fatal(err)
	}
}
