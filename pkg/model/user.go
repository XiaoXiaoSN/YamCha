package model

import "time"

// User Object
type User struct {
	ID        int       `gorm:"id" json:"id"`
	Name      string    `gorm:"name" json:"name" validate:"required"`
	LineID    string    `gorm:"line_id" json:"line_id" validate:"required"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
