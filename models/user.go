package models

import "gorm.io/gorm"

type Role int

const (
	RoleAdmin Role = iota
	RoleUser
)

type User struct {
	gorm.Model

	TelegramID int64 `gorm:"unique"`
	Role
}

func NewUser(telegramID int64) *User {
	return &User{
		TelegramID: telegramID,
	}
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}
