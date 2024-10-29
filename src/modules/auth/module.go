package auth

import (
	"github.com/thmelodev/ddd-events-api/src/modules/auth/application/usecases"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/domain/repositories"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/infra/mappers"
	infraRepositories "github.com/thmelodev/ddd-events-api/src/modules/auth/infra/repositories"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"auth",
		fx.Provide(usecases.Usecases...),
		fx.Provide(mappers.NewUserMapper),
		fx.Provide(
			fx.Annotate(infraRepositories.NewUserRepository, fx.As(new(repositories.IUserRepository))),
		),
		fx.Invoke(NewAuthController),
	)
}
