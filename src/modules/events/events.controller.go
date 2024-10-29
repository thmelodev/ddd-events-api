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
	conf                  *config.Config
	httpServer            *httpServer.HttpServer
	createEventUsecase    *usecases.CreateEventUsecase
	deleteEventUsecase    *usecases.DeleteEventUsecase
	updateEventUsecase    *usecases.UpdateEventUsecase
	getEventsQuery        *queries.GetEventsQuery
	getEventByIdQuery     *queries.GetEventByIdQuery
	getEventByUserIdQuery *queries.GetEventByUserIdQuery
}

func NewEventsController(
	conf *config.Config,
	hs *httpServer.HttpServer,
	createEventUsecase *usecases.CreateEventUsecase,
	deleteEventUsecase *usecases.DeleteEventUsecase,
	updateEventUsecase *usecases.UpdateEventUsecase,
	getEventsQuery *queries.GetEventsQuery,
	getEventByIdQuery *queries.GetEventByIdQuery,
	getEventByUserIdQuery *queries.GetEventByUserIdQuery,
) *EventsController {

	httpGroup := hs.AppGroup.Group("/events")
	httpGroup.Use(httpServer.ErrorHandler())

	controller := &EventsController{
		conf:                  conf,
		httpServer:            hs,
		createEventUsecase:    createEventUsecase,
		deleteEventUsecase:    deleteEventUsecase,
		updateEventUsecase:    updateEventUsecase,
		getEventsQuery:        getEventsQuery,
		getEventByIdQuery:     getEventByIdQuery,
		getEventByUserIdQuery: getEventByUserIdQuery,
	}

	controller.registerRoutes(httpGroup)

	return controller
}

func (ec *EventsController) registerRoutes(group *gin.RouterGroup) {
	group.GET("/health", ec.health)
	group.GET("", ec.getEvents)
	group.GET("/:id", ec.getEventById)
	group.GET("/user/:userId", ec.getEventByUserId)
	group.POST("", httpServer.AuthenticationHandler(ec.conf), ec.createEvent)
	group.PUT("/:id", httpServer.AuthenticationHandler(ec.conf), ec.updateEvent)
	group.DELETE("/:id", httpServer.AuthenticationHandler(ec.conf), ec.deleteEvent)
}

func (ec *EventsController) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func (ec *EventsController) getEvents(ctx *gin.Context) {
	result, err := ec.getEventsQuery.Execute(ctx, nil)

	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, result)
}

func (ec *EventsController) createEvent(ctx *gin.Context) {
	var event usecases.CreateEventUsecaseProps
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.Error(apiErrors.NewInvalidPropsError(err.Error()))
		return
	}

	event.UserId = ctx.Request.Header.Get("X-User-Id")

	result, err := ec.createEventUsecase.Execute(ctx, &event)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (ec *EventsController) getEventById(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := ec.getEventByIdQuery.Execute(ctx, id)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (ec *EventsController) getEventByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")

	result, err := ec.getEventByUserIdQuery.Execute(ctx, userId)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (ec *EventsController) deleteEvent(ctx *gin.Context) {
	var event usecases.DeleteEventProps
	event.Id = ctx.Param("id")
	event.UserId = ctx.Request.Header.Get("X-User-Id")

	result, err := ec.deleteEventUsecase.Execute(ctx, &event)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (ec *EventsController) updateEvent(ctx *gin.Context) {
	id := ctx.Param("id")

	var event usecases.UpdateEventDTO

	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.Error(apiErrors.NewInvalidPropsError(err.Error()))
		return
	}

	event.Id = id
	event.UserId = ctx.Request.Header.Get("X-User-Id")

	result, err := ec.updateEventUsecase.Execute(ctx, &event)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
