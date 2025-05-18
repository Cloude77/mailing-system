package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	//Инициализируем логгер
	log := logrus.New()
	log.Info("Starting API Service")

	// Создаём Gin-роутер
	r := gin.Default()

	// Тестовый	 эндпоинт
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Запускаем сервер
	if err := r.Run(":8083"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
