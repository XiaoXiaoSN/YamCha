//+build wireinject

package line

import (
	"context"

	"yamcha/internal/config"
	"yamcha/internal/httputil"
	"yamcha/internal/provider"
	"yamcha/pkg/delivery/api"
	"yamcha/pkg/delivery/linebot"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

type Application struct {
	Echo       *echo.Echo
	Controller api.Controller
	LineBot    linebot.LineBot
}

// InitApplication ...
func InitApplication(ctx context.Context) (*Application, error) {
	wire.Build(
		config.NewConfiguration,
		wire.FieldsOf(new(*config.Configuration), "DBCfg", "BotCfg"),

		httputil.NewEcho,

		provider.GORMSet,
		provider.RepoSet,
		provider.ServiceSet,

		// delivery
		api.NewController,
		linebot.NewYambotLineBot,

		wire.Struct(new(Application), "*"),
	)
	return nil, nil
}
