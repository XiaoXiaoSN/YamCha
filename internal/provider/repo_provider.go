package provider

import (
	"yamcha/pkg/repository/db"

	"github.com/google/wire"
)

// RepoSet ...
var RepoSet = wire.NewSet(
	db.NewRepo,
)
