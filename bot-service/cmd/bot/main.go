package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Инициализируем логгер
	log := logrus.New()
	log.Info("Starting Bot Service")

	// Читаем конфигурацию
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read .env: %v", err)
	}

	// Получаем Telegram-токен
	token := viper.GetString("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN is not set")
	}

	// Создаём Telegram-бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	log.Infof("Authorized on account %s", bot.Self.UserName)

	// Тестовый запуск
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Infof("Received message: %s", update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Echo: "+update.Message.Text)
		bot.Send(msg)
	}
}
