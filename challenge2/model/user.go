package model

import (
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Role      string `gorm:"type:varchar(5)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
