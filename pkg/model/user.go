package model

import "time"

// User Object
type User struct {
	ID        int       `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name" validate:"required"`
	LineID    string    `gorm:"column:line_id" json:"line_id" validate:"required"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
