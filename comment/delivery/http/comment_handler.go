package delivery

import (
	"fmt"
	"mygram-api/comment/delivery/http/middleware"
	"mygram-api/comment/utils"
	"mygram-api/domain"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type commentRoute struct {
	commentUseCase domain.CommentUseCase
	photoUseCase   domain.PhotoUseCase
}

func NewCommentRoute(handlers *gin.Engine, commentUseCase domain.CommentUseCase, photoUseCase domain.PhotoUseCase) {
	route := &commentRoute{commentUseCase, photoUseCase}

	handler := handlers.Group("/comments")
	{
		handler.Use(middleware.Authentication())
		handler.GET("", route.Fetch)
		handler.POST("", route.Store)
		handler.PUT("/:commentId", middleware.Authorization(route.commentUseCase), route.Update)
		handler.DELETE("/:commentId", middleware.Authorization(route.commentUseCase), route.Delete)
	}
}

func (route *commentRoute) Fetch(ctx *gin.Context) {
	var (
		comments []domain.Comment

		err error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = route.commentUseCase.Fetch(ctx.Request.Context(), &comments, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (route *commentRoute) Store(ctx *gin.Context) {
	var (
		comment domain.Comment
		photo   domain.Photo
		err     error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	photoID := comment.PhotoID

	if err = route.photoUseCase.GetByID(ctx.Request.Context(), &photo, photoID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Not Found",
			"message": fmt.Sprintf("photo with id %s doesn't exist", photoID),
		})

		return
	}

	comment.UserID = userID

	if err = route.commentUseCase.Store(ctx.Request.Context(), &comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, utils.NewComment{
		ID:        comment.ID,
		UserID:    comment.UserID,
		PhotoID:   comment.PhotoID,
		Message:   comment.Message,
		CreatedAt: comment.CreatedAt,
	})
}

func (route *commentRoute) Update(ctx *gin.Context) {
	var (
		comment domain.Comment
		photo   domain.Photo
		err     error
	)

	commentID := ctx.Param("commentId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	updatedComment := domain.Comment{
		UserID:  userID,
		Message: comment.Message,
	}

	if photo, err = route.commentUseCase.Update(ctx.Request.Context(), updatedComment, commentID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, utils.UpdatedComment{
		ID:        photo.ID,
		UserID:    photo.UserID,
		Title:     photo.Title,
		PhotoUrl:  photo.PhotoUrl,
		Caption:   photo.Caption,
		UpdatedAt: photo.UpdatedAt,
	})
}

func (route *commentRoute) Delete(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	if err := route.commentUseCase.Delete(ctx.Request.Context(), commentID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})
}
