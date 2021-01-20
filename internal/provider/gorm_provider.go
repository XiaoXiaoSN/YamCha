package provider

import (
	"yamcha/internal/database"

	"github.com/google/wire"
)

// GORMSet 初始化 gorm database
var GORMSet = wire.NewSet(
	database.NewDatabases,
)
