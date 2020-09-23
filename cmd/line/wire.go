//+build wireinject

package line

import (
	"context"

	"yamcha/internal/config"
	"yamcha/internal/provider"
	"yamcha/pkg/repository"

	"github.com/google/wire"
)

func initApplication(ctx context.Context) (repository.Repository, error) {
	wire.Build(
		config.NewConfiguration,
		provider.GORMSet,
		provider.RepoSet,

		repository.NewRepo,
	)
	return nil, nil
}
