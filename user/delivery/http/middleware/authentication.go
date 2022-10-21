package middleware

import (
	"mygram-api/helpers"
	"mygram-api/user/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseMessage{
				Status:  "unauthenticated",
				Message: err.Error(),
			})

			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
