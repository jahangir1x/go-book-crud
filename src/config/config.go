package config

import (
	"app/src/logger"

	"github.com/spf13/viper"
)

type App struct {
	Name     string
	Host     string
	Port     string
	Version  string
	LogLevel string
}

type Database struct {
	Provider       string
	Host           string
	Port           string
	User           string
	Password       string
	Name           string
	SslMode        string
	Timezone       string
	MaxConnections int
}

type Config struct {
	App      App
	Database Database
}

var appConfig = Config{}

func InitConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("Error reading config file", err)
	}

	appConfig = Config{
		App: App{
			Name:     viper.GetString("APP_NAME"),
			Host:     viper.GetString("APP_HOST"),
			Port:     viper.GetString("APP_PORT"),
			Version:  viper.GetString("APP_VERSION"),
			LogLevel: viper.GetString("APP_LOG_LEVEL"),
		},
		Database: Database{
			Provider:       viper.GetString("DB_PROVIDER"),
			Host:           viper.GetString("DB_HOST"),
			Port:           viper.GetString("DB_PORT"),
			User:           viper.GetString("DB_USER"),
			Password:       viper.GetString("DB_PASSWORD"),
			Name:           viper.GetString("DB_NAME"),
			SslMode:        viper.GetString("DB_SSL_MODE"),
			Timezone:       viper.GetString("DB_TIMEZONE"),
			MaxConnections: viper.GetInt("DB_MAX_CONNECTIONS"),
		},
	}
}

func GetConfig() Config {
	return appConfig
}
