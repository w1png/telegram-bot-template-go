package sqlite

import (
	"reflect"

	"github.com/w1png/telegram-bot-template/errors"
	"github.com/w1png/telegram-bot-template/models"
	"gorm.io/gorm"
)

func (s *SQLiteStorage) CreateUser(user *models.User) error {
	if err := s.DB.Create(user).Error; err != nil && err.Error() == "UNIQUE constraint failed: users.telegram_id" {
		return errors.NewObjectAlreadyExistsError("User")
	} else if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteStorage) GetUserByTelegramID(telegramID int64) (*models.User, error) {
	var user models.User

	if err := s.DB.Where("telegram_id = ?", telegramID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewObjectNotFoundError("User")
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (s *SQLiteStorage) GetUsers(offset int, limit int) ([]models.User, error) {
	var users []models.User
	err := s.DB.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (s *SQLiteStorage) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := s.DB.Find(&users).Error
	return users, err
}

func (s *SQLiteStorage) UpdateUser(user *models.User) error {
	if _, err := s.GetUserByTelegramID(user.TelegramID); err != nil {
		return err
	}

	return s.DB.Save(user).Error
}

func (s *SQLiteStorage) DeleteUser(telegramID int64) error {
	if _, err := s.GetUserByTelegramID(telegramID); err != nil {
		return err
	}

	return s.DB.Where("telegram_id = ?", telegramID).Delete(&models.User{}).Error
}

func (s *SQLiteStorage) CountUsers() (int, error) {
	var count int64
	err := s.DB.Model(&models.User{}).Count(&count).Error
	return int(count), err
}

func (s *SQLiteStorage) CreateUserIfDoesntExist(telegram_id int64) error {
	if _, err := s.GetUserByTelegramID(telegram_id); err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(errors.NewObjectNotFoundError("User")) {
			user := models.NewUser(telegram_id)
			if count, err := s.CountUsers(); err != nil {
				return err
			} else {
				if count == 0 {
					user.Role = models.RoleAdmin
				}
			}
			return s.CreateUser(user)
		} else {
			return err
		}
	}

	return nil
}
