package events

import (
	applicationMapper "github.com/thmelodev/ddd-events-api/src/modules/events/application/mappers"
	"github.com/thmelodev/ddd-events-api/src/modules/events/application/queries"
	"github.com/thmelodev/ddd-events-api/src/modules/events/application/usecases"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain/repositories"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/mappers"
	infraRepositories "github.com/thmelodev/ddd-events-api/src/modules/events/infra/repositories"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"events",
		fx.Provide(
			fx.Annotate(infraRepositories.NewEventRepository, fx.As(new(repositories.IEventRepository))),
		),
		fx.Provide(applicationMapper.NewEventMapper),
		fx.Provide(mappers.NewEventMapper),
		fx.Provide(usecases.Usecases...),
		fx.Provide(queries.Queries...),
		fx.Invoke(NewEventsController),
	)
}
