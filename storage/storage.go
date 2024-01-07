package storage

import (
	"github.com/w1png/telegram-bot-template/config"
	"github.com/w1png/telegram-bot-template/errors"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/storage/sqlite"
)

var StorageInstance Storage

type Storage interface {
	CreateUser(user *models.User) error
	GetUserByTelegramID(telegramID int64) (*models.User, error)
	GetUsers(offset int, limit int) ([]models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(telegramID int64) error
	CountUsers() (int, error)
	CreateUserIfDoesntExist(telegram_id int64) error
}

func InitStorage() error {
	var err error

	switch config.ConfigInstance.StorageType {
	case config.SQLite:
		{
			StorageInstance, err = sqlite.NewSQLiteStorage()
		}
	}

	if err != nil {
		return err
	}

	if StorageInstance == nil {
		return errors.NewUnknownStorageTypeError(string(config.ConfigInstance.StorageType))
	}

	return nil
}
