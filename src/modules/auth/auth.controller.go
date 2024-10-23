package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/application/usecases"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/providers/httpServer"
)

type AuthController struct {
	conf              *config.Config
	httpServer        *httpServer.HttpServer
	createUserUsecase *usecases.CreateUserUsecase
}

func NewAuthController(
	conf *config.Config,
	hs *httpServer.HttpServer,
	createUserUsecase *usecases.CreateUserUsecase,
) *AuthController {

	httpGroup := hs.AppGroup.Group("/auth")
	httpGroup.Use(httpServer.ErrorHandler())

	hs.AppGroup = httpGroup

	controller := &AuthController{
		conf:              conf,
		httpServer:        hs,
		createUserUsecase: createUserUsecase,
	}

	controller.registerRoutes()

	return controller
}

func (ac *AuthController) registerRoutes() {
	ac.httpServer.AppGroup.POST("/register", ac.register)
	// ac.httpServer.AppGroup.POST("/login", ac.login)
}

func (ac *AuthController) register(ctx *gin.Context) {
	result, err := ac.createUserUsecase.Execute(ctx, nil)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
