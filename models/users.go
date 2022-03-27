package models

import "time"

type UsersResponse struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `gorm:"created_at"`
}
type CreateAccountRequest struct {
	UserName string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	User  UsersResponse `json:"user"`
	Token string        `json:"token"`
}
