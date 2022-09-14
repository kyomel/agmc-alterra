package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
