package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique_index;not null"`
	PasswordHash string `gorm:"not null"`
	HasUsedFree  bool   `gorm:"default:false"`
	Role         string `gorm:"type:varchar(20);default:'user'"`
}
