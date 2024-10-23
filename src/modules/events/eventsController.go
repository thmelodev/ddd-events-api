package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/modules/events/application/queries"
	"github.com/thmelodev/ddd-events-api/src/modules/events/application/usecases"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/providers/httpServer"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

type EventsController struct {
	conf               *config.Config
	httpServer         *httpServer.HttpServer
	createEventUsecase *usecases.CreateEventUsecase
	deleteEventUsecase *usecases.DeleteEventUsecase
	updateEventUsecase *usecases.UpdateEventUsecase
	getEventsQuery     *queries.GetEventsQuery
	getEventByIdQuery  *queries.GetEventByIdQuery
}

func NewEventsController(
	conf *config.Config,
	hs *httpServer.HttpServer,
	createEventUsecase *usecases.CreateEventUsecase,
	deleteEventUsecase *usecases.DeleteEventUsecase,
	updateEventUsecase *usecases.UpdateEventUsecase,
	getEventsQuery *queries.GetEventsQuery,
	getEventByIdQuery *queries.GetEventByIdQuery,
) *EventsController {

	httpGroup := hs.AppGroup.Group("/events")
	httpGroup.Use(httpServer.ErrorHandler())

	hs.AppGroup = httpGroup

	controller := &EventsController{
		conf:               conf,
		httpServer:         hs,
		createEventUsecase: createEventUsecase,
		deleteEventUsecase: deleteEventUsecase,
		updateEventUsecase: updateEventUsecase,
		getEventsQuery:     getEventsQuery,
		getEventByIdQuery:  getEventByIdQuery,
	}

	controller.registerRoutes()

	return controller
}

func (ec EventsController) registerRoutes() {
	ec.httpServer.AppGroup.GET("/health", ec.health)
	ec.httpServer.AppGroup.GET("", ec.getEvents)
	ec.httpServer.AppGroup.GET("/:id", ec.getEventById)
	ec.httpServer.AppGroup.POST("", ec.createEvent)
	ec.httpServer.AppGroup.PUT("/:id", ec.updateEvent)
	ec.httpServer.AppGroup.DELETE("/:id", ec.deleteEvent)

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
	var event usecases.CreateEventDTO
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.Error(apiErrors.NewInvalidPropsError(err.Error()))
		return
	}

	result, err := ec.createEventUsecase.Execute(ctx, &event)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(201, result)
}

func (ec EventsController) getEventById(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := ec.getEventByIdQuery.Execute(ctx, id)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (ec EventsController) deleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := ec.deleteEventUsecase.Execute(ctx, id)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (ec EventsController) updateEvent(ctx *gin.Context) {
	id := ctx.Param("id")

	var event usecases.UpdateEventDTO

	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.Error(apiErrors.NewInvalidPropsError(err.Error()))
		return
	}

	event.Id = id

	result, err := ec.updateEventUsecase.Execute(ctx, &event)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
