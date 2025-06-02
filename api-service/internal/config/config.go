package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config хранит настройки приложения
type Config struct {
	APIHost       string `mapstructure:"API_HOST"`
	APIPort       string `mapstructure:"API_PORT"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_Password"`
	DBName        string `mapstructure:"DB_NAME"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPORT     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	GRPCHost      string `mapstructure:"GRPC_BOT_HOST"`
	GRPCPort      string `mapstructure:"GRPC_BOT_PORT"`
}

// Load загружает конфигурацию из .env
func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("Failed to read .env: %v", err)
		return nil, err
	}
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		logrus.Errorf("Failed to unmatshal config: %v", err)
		return nil, err
	}
	return cfg, nil
}
