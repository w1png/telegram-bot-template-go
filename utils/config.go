package utils

import (
	"os"

	"github.com/w1png/telegram-bot-template/errors"
)

type Config struct {
  TelegramToken string
  Language string
  StorageType string
  LoggerType string
}

func (c *Config) GatherVariables() error {
  token, ok := os.LookupEnv("TELEGRAM_TOKEN")
  if !ok {
    return errors.NewEnvironmentVariableError("TELEGRAM_TOKEN")
  }

  loggerType, ok := os.LookupEnv("LOGGER_TYPE")
  if !ok {
    loggerType = ""
  }

  language, ok := os.LookupEnv("LANGUAGE")
  if !ok {
    return errors.NewEnvironmentVariableError("LANGUAGE")
  }

  storageType, ok := os.LookupEnv("STORAGE_TYPE")
  if !ok {
    return errors.NewEnvironmentVariableError("STORAGE_TYPE")
  }

  c.TelegramToken = token
  c.Language = language
  c.StorageType = storageType
  c.LoggerType = loggerType

  return nil
}
