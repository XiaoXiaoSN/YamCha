package user

import (
	"context"
	"time"
)

// User Object
type User struct {
	ID        int       `gorm:"id" json:"id"`
	Name      string    `gorm:"name" json:"name"`
	LineID    string    `gorm:"line_id" json:"line_id"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

// Service is a user service
type Service interface {
	CreateUser(ctx context.Context, u User) error
}

// Repository is a user repo
type Repository interface {
	CreateUser(ctx context.Context, u User) error
}
