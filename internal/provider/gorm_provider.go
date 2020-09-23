package provider

import (
	"yamcha/internal/config"
	"yamcha/internal/database"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// GORMSet ...
var GORMSet = wire.NewSet(
	InitGORM,
)

// InitGORM 初始化 gorm database
func InitGORM(cfg *config.Configuration) (*gorm.DB, error) {
	return database.NewDatabases(cfg.DBCfg)
}
