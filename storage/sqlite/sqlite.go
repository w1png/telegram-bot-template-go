package sqlite

import (
	"github.com/w1png/telegram-bot-template/config"
	"github.com/w1png/telegram-bot-template/errors"
	"github.com/w1png/telegram-bot-template/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLiteStorage struct {
	DB *gorm.DB
}

func NewSQLiteStorage() (*SQLiteStorage, error) {
	storage := &SQLiteStorage{}

	var err error
	if storage.DB, err = gorm.Open(sqlite.Open(config.ConfigInstance.SQLitePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}); err != nil {
		return nil, errors.NewDatabaseConnectionError(err.Error())
	}

	err = storage.autoMigrate()
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *SQLiteStorage) autoMigrate() error {
	err := s.DB.AutoMigrate(&models.User{})
	if err != nil {
		return errors.NewDatabaseMigrationError(err.Error())
	}

	return nil
}
