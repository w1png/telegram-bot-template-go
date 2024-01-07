package config

import (
	"os"

	"github.com/w1png/telegram-bot-template/errors"
)

var ConfigInstance *Config

type StorageType string

const (
	SQLite StorageType = "sqlite"
)

type Config struct {
	TelegramToken string
	IsDebug       bool
	Language      string
	StorageType   StorageType
	SQLitePath    string
	LoggerType    string
}

func InitConfig() error {
	ConfigInstance = &Config{}

	return ConfigInstance.GatherVariables()
}

func (c *Config) GatherVariables() error {
	var ok bool
	c.TelegramToken, ok = os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		return errors.NewEnvironmentVariableError("TELEGRAM_TOKEN")
	}

	c.IsDebug = os.Getenv("DEBUG") == "true"

	c.LoggerType, ok = os.LookupEnv("LOGGER_TYPE")
	if !ok {
		c.LoggerType = ""
	}

	c.Language, ok = os.LookupEnv("LANGUAGE")
	if !ok {
		return errors.NewEnvironmentVariableError("LANGUAGE")
	}

	storage_type, ok := os.LookupEnv("STORAGE_TYPE")
	if !ok {
		return errors.NewEnvironmentVariableError("STORAGE_TYPE")
	}

	switch storage_type {
	case string(SQLite):
		path, ok := os.LookupEnv("SQLITE_PATH")
		if !ok {
			return errors.NewEnvironmentVariableError("SQLITE_PATH")
		}
		c.SQLitePath = path
	default:
		return errors.NewEnvironmentVariableError("STORAGE_TYPE")
	}

	c.StorageType = StorageType(storage_type)

	return nil
}
