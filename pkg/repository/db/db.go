package db

import (
	"yamcha/pkg/repository"

	"github.com/jinzhu/gorm"
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
