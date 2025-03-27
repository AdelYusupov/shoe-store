package models

import "gorm.io/gorm"

type AdminUser struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
}
