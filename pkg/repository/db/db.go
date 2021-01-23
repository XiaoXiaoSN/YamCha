package db

import (
	"yamcha/pkg/repository"

	"gorm.io/gorm"
)

type dbRepository struct {
	db *gorm.DB
}

// NewRepo create a database object who implement the repository
func NewRepo(db *gorm.DB) repository.Repository {
	return &dbRepository{
		db: db,
	}
}
