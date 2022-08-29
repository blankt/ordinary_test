package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model

	Username string `json:"username"`
	Password string `json:"password"`
}
