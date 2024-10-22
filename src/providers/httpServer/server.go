package httpServer

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/utils/logger"
	"go.uber.org/fx"
)

type HttpServer struct {
	AppGroup  *gin.RouterGroup
	AppServer *gin.Engine
}

func NewServer(
	config *config.Config,
	lc fx.Lifecycle,
) *HttpServer {
	log := logger.Get()

	server := gin.Default()

	appGroup := server.Group(fmt.Sprintf("/%s", os.Getenv("APP_NAME")))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Http.HttpPort),
		Handler: server,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Debugf("Starting server in port: %d", config.Http.HttpPort)
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Errorf("Error when starting server on port %d: %s", config.Http.HttpPort, err)
				}
			}()
			return nil

		},
		OnStop: func(ctx context.Context) error {
			log.Debugf("Stopping the server...")
			shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(shutdownCtx); err != nil {
				log.Errorf("Server shutdown failed: %s", err)
				return err
			}
			log.Debugf("Server stopped gracefully")
			return nil
		},
	})

	return &HttpServer{
		AppGroup:  appGroup,
		AppServer: server,
	}

}
