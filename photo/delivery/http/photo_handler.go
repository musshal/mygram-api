package delivery

import (
	"mygram-api/domain"
	"mygram-api/photo/delivery/http/middleware"
	"mygram-api/photo/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type photoRoute struct {
	photoUseCase domain.PhotoUseCase
}

func NewPhotoRoute(handlers *gin.Engine, photoUseCase domain.PhotoUseCase) {
	route := &photoRoute{photoUseCase}

	handler := handlers.Group("/photos")
	{
		handler.Use(middleware.Authentication())
		handler.GET("", route.Fetch)
		handler.POST("", route.Store)
		handler.PUT("/:photoId", middleware.Authorization(route.photoUseCase), route.Update)
		handler.DELETE("/:photoId", middleware.Authorization(route.photoUseCase), route.Delete)
	}
}

func (route *photoRoute) Fetch(ctx *gin.Context) {
	var (
		photos []domain.Photo
		err    error
	)

	if err = route.photoUseCase.Fetch(ctx.Request.Context(), &photos); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	fetchedPhotos := []*utils.Photo{}

	for _, photo := range photos {
		fetchedPhotos = append(fetchedPhotos, &utils.Photo{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: &utils.User{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		})
	}

	ctx.JSON(http.StatusOK, fetchedPhotos)
}

func (route *photoRoute) Store(ctx *gin.Context) {
	var (
		photo domain.Photo
		err   error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	photo.UserID = userID

	if err = route.photoUseCase.Store(ctx.Request.Context(), &photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, utils.NewPhoto{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	})
}

func (route *photoRoute) Update(ctx *gin.Context) {
	var (
		photo domain.Photo
		err   error
	)

	if err = ctx.ShouldBindJSON(&photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	updatedPhoto := domain.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}

	photoID := ctx.Param("photoId")

	if photo, err = route.photoUseCase.Update(ctx.Request.Context(), updatedPhoto, photoID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, utils.UpdatedPhoto{
		ID:        photo.ID,
		UserID:    photo.UserID,
		Title:     photo.Title,
		PhotoUrl:  photo.PhotoUrl,
		Caption:   photo.Caption,
		UpdatedAt: photo.UpdatedAt,
	})
}

func (route *photoRoute) Delete(ctx *gin.Context) {
	photoID := ctx.Param("photoId")

	if err := route.photoUseCase.Delete(ctx.Request.Context(), photoID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your photo has been successfully deleted",
	})
}
