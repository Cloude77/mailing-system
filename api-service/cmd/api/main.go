package main

import (
	"fmt"
	"github.com/Cloude77/mailing-system/api-service/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	//Инициализируем логгер
	log := logrus.New()
	log.Info("Starting API Service")

	// Загружаем конфигурацию
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Создаём Gin-роутер
	r := gin.Default()
	// Тестовый	 эндпоинт
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Запускаем сервер
	addr := fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
