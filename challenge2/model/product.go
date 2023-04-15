package model

import "time"

type Product struct {
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255)"`
	Price     int
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductCreateRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
type ProductCreateResponse struct {
	ID        string    `json:"product_id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductResponse struct {
	ID        string    `json:"product_id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type ProductUpdateRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
