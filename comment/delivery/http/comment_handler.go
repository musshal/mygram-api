package delivery

import (
	"fmt"
	"mygram-api/comment/delivery/http/middleware"
	"mygram-api/comment/utils"
	"mygram-api/domain"
	"mygram-api/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentUseCase domain.CommentUseCase
	photoUseCase   domain.PhotoUseCase
}

func NewCommentHandler(routers *gin.Engine, commentUseCase domain.CommentUseCase, photoUseCase domain.PhotoUseCase) {
	handler := &commentHandler{commentUseCase, photoUseCase}

	router := routers.Group("/comments")
	{
		router.Use(middleware.Authentication())
		router.GET("", handler.Fetch)
		router.POST("", handler.Store)
		router.PUT("/:commentId", middleware.Authorization(handler.commentUseCase), handler.Update)
		router.DELETE("/:commentId", middleware.Authorization(handler.commentUseCase), handler.Delete)
	}
}

func (handler *commentHandler) Fetch(ctx *gin.Context) {
	var (
		comments []domain.Comment

		err error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = handler.commentUseCase.Fetch(ctx.Request.Context(), &comments, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data:   comments,
	})
}

func (handler *commentHandler) Store(ctx *gin.Context) {
	var (
		comment domain.Comment
		photo   domain.Photo
		err     error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	photoID := comment.PhotoID

	if err = handler.photoUseCase.GetByID(ctx.Request.Context(), &photo, photoID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.ResponseMessage{
			Status:  "fail",
			Message: fmt.Sprintf("photo with id %s doesn't exist", photoID),
		})

		return
	}

	comment.UserID = userID

	if err = handler.commentUseCase.Store(ctx.Request.Context(), &comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, helpers.ResponseData{
		Status: "success",
		Data: utils.NewComment{
			ID:        comment.ID,
			UserID:    comment.UserID,
			PhotoID:   comment.PhotoID,
			Message:   comment.Message,
			CreatedAt: comment.CreatedAt,
		},
	})
}

func (handler *commentHandler) Update(ctx *gin.Context) {
	var (
		comment domain.Comment
		photo   domain.Photo
		err     error
	)

	commentID := ctx.Param("commentId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	updatedComment := domain.Comment{
		UserID:  userID,
		Message: comment.Message,
	}

	if photo, err = handler.commentUseCase.Update(ctx.Request.Context(), updatedComment, commentID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.ResponseData{
		Status: "success",
		Data: utils.UpdatedComment{
			ID:        photo.ID,
			UserID:    photo.UserID,
			Title:     photo.Title,
			PhotoUrl:  photo.PhotoUrl,
			Caption:   photo.Caption,
			UpdatedAt: photo.UpdatedAt,
		},
	})
}

func (handler *commentHandler) Delete(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	if err := handler.commentUseCase.Delete(ctx.Request.Context(), commentID); err != nil {
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
