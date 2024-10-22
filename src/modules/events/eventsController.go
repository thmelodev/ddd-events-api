package events

import (
	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/modules/events/application/queries"
	"github.com/thmelodev/ddd-events-api/src/modules/events/application/usecase"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/models"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/providers/httpServer"
)

type EventsController struct {
	conf               *config.Config
	httpServer         *httpServer.HttpServer
	createEventUsecase *usecase.CreateEventUsecase
	getEventsQuery     *queries.GetEventsQuery
}

func NewEventsController(
	conf *config.Config,
	httpServer *httpServer.HttpServer,
	createEventUsecase *usecase.CreateEventUsecase,
	getEventsQuery *queries.GetEventsQuery,
) *EventsController {

	httpGroup := httpServer.AppGroup.Group("/events")
	httpServer.AppGroup = httpGroup

	controller := &EventsController{
		conf:               conf,
		httpServer:         httpServer,
		createEventUsecase: createEventUsecase,
		getEventsQuery:     getEventsQuery,
	}

	controller.registerRoutes()

	return controller
}

func (ec EventsController) registerRoutes() {
	ec.httpServer.AppGroup.GET("/health", ec.health)
	ec.httpServer.AppGroup.GET("", ec.getEvents)
	ec.httpServer.AppGroup.POST("", ec.createEvent)
}

func (ec EventsController) health(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (ec EventsController) getEvents(ctx *gin.Context) {
	result, err := ec.getEventsQuery.Execute(ctx, nil)

	if err != nil {
		ctx.JSON(400, err)
	}

	ctx.JSON(200, result)
}

func (ec EventsController) createEvent(ctx *gin.Context) {
	var event models.Event
	ctx.ShouldBindJSON(&event)

	result, err := ec.createEventUsecase.Execute(ctx, nil)

	if err != nil {
		ctx.JSON(400, err)
	}

	ctx.JSON(200, result)
}
