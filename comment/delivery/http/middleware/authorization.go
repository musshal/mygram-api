package middleware

import (
	"fmt"
	"mygram-api/domain"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization(commentUseCase domain.CommentUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			comment domain.Comment
			err     error
		)

		commentID := ctx.Param("commentId")
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := string(userData["id"].(string))

		if err = commentUseCase.GetByID(ctx.Request.Context(), &comment, commentID); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": fmt.Sprintf("comment with id %s doesn't exist", commentID),
			})

			return
		}

		if comment.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you don't have permission to view or edit this comment",
			})

			return
		}
	}
}
