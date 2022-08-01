package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password []byte `json:"-"`
}
