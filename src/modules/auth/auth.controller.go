package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/application/usecases"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/providers/httpServer"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

type AuthController struct {
	conf              *config.Config
	httpServer        *httpServer.HttpServer
	createUserUsecase *usecases.CreateUserUsecase
	loginUserUsecase  *usecases.LoginUserUsecase
}

func NewAuthController(
	conf *config.Config,
	hs *httpServer.HttpServer,
	createUserUsecase *usecases.CreateUserUsecase,
	loginUserUsecase *usecases.LoginUserUsecase,
) *AuthController {

	httpGroup := hs.AppGroup.Group("/auth")
	httpGroup.Use(httpServer.ErrorHandler())

	controller := &AuthController{
		conf:              conf,
		httpServer:        hs,
		createUserUsecase: createUserUsecase,
		loginUserUsecase:  loginUserUsecase,
	}

	controller.registerRoutes(httpGroup)

	return controller
}

func (ac *AuthController) registerRoutes(group *gin.RouterGroup) {
	group.POST("/register", httpServer.AuthenticationHandler(ac.conf), ac.register)
	group.POST("/login", ac.login)
}

func (ac *AuthController) register(ctx *gin.Context) {
	var userProps usecases.CreateUserDTO
	if err := ctx.ShouldBindJSON(&userProps); err != nil {
		ctx.Error(apiErrors.NewInvalidPropsError(err.Error()))
		return
	}

	userProps.UserId = ctx.Request.Header.Get("X-User-Id")

	result, err := ac.createUserUsecase.Execute(ctx, &userProps)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (ac *AuthController) login(ctx *gin.Context) {
	var userProps usecases.LoginUserDTO
	if err := ctx.ShouldBindJSON(&userProps); err != nil {
		ctx.Error(apiErrors.NewInvalidPropsError(err.Error()))
		return
	}

	result, err := ac.loginUserUsecase.Execute(ctx, &userProps)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
