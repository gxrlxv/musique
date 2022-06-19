package domain

import (
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Gender       string    `json:"gender"`
	Country      string    `json:"country"`
	City         string    `json:"city"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at"`
	Role         int       `json:"role"`
}

type CreateUserDTO struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"-"`
	RepeatPassword string `json:"repeat_password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Gender         string `json:"gender"`
	Country        string `json:"country"`
	City           string `json:"city"`
	Phone          string `json:"phone"`
}

func NewUser(dto CreateUserDTO, passwordHash string) *User {
	return &User{
		ID:           "",
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: passwordHash,
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		Gender:       dto.Gender,
		Country:      dto.Country,
		City:         dto.City,
		Phone:        dto.Phone,
		CreatedAt:    time.Now(),
		Role:         0,
	}
}
