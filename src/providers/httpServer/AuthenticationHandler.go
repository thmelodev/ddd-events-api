package httpServer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/utils/jwt"
)

func AuthenticationHandler(cf *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "missing token",
			})
			ctx.Abort()
			return
		}

		tokenClaims, err := jwt.ValidateToken(token, cf)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Request.Header.Set("X-User-Id", tokenClaims.UserID)
		ctx.Request.Header.Set("X-User-Email", tokenClaims.Email)

		ctx.Next()
	}
}
