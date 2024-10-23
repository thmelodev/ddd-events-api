package auth

import (
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/providers/httpServer"
)

type AuthController struct {
	conf       *config.Config
	httpServer *httpServer.HttpServer
}

func NewAuthController(
	conf *config.Config,
	hs *httpServer.HttpServer,
) *AuthController {

	httpGroup := hs.AppGroup.Group("/auth")
	httpGroup.Use(httpServer.ErrorHandler())

	hs.AppGroup = httpGroup

	controller := &AuthController{
		conf:       conf,
		httpServer: hs,
	}

	controller.registerRoutes()

	return controller
}

func (ac *AuthController) registerRoutes() {

}
