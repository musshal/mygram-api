package delivery

import (
	"mygram-api/domain"
	"mygram-api/socialmedia/delivery/http/middleware"
	"mygram-api/socialmedia/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type socialMediaRoute struct {
	socialMediaUseCase domain.SocialMediaUseCase
}

func NewSocialMediaRoute(handlers *gin.Engine, socialMediaUseCase domain.SocialMediaUseCase) {
	route := &socialMediaRoute{socialMediaUseCase}

	handler := handlers.Group("/socialmedias")
	{
		handler.Use(middleware.Authentication())
		handler.GET("", route.Fetch)
		handler.POST("", route.Store)
		handler.PUT("/:socialMediaId", middleware.Authorization(route.socialMediaUseCase), route.Update)
		handler.DELETE("/:socialMediaId", middleware.Authorization(route.socialMediaUseCase), route.Delete)
	}
}

func (route *socialMediaRoute) Fetch(ctx *gin.Context) {
	var (
		socialMedias []domain.SocialMedia
		err          error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = route.socialMediaUseCase.Fetch(ctx.Request.Context(), &socialMedias, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"socialMedia": socialMedias,
	})
}
func (route *socialMediaRoute) Store(ctx *gin.Context) {
	var (
		socialMedia domain.SocialMedia
		err         error
	)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	socialMedia.UserID = userID

	if err = route.socialMediaUseCase.Store(ctx.Request.Context(), &socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, utils.NewSocialMedia{
		ID:             socialMedia.ID,
		UserID:         socialMedia.UserID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		CreatedAt:      socialMedia.CreatedAt,
	})
}

func (route *socialMediaRoute) Update(ctx *gin.Context) {
	var (
		socialMedia domain.SocialMedia
		err         error
	)

	socialMediaID := ctx.Param("socialMediaId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	if err = ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	updatedSocialMedia := domain.SocialMedia{
		UserID:         userID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}

	if socialMedia, err = route.socialMediaUseCase.Update(ctx.Request.Context(), updatedSocialMedia, socialMediaID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, utils.UpdatedSocialMedia{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		UpdatedAt:      socialMedia.UpdatedAt,
	})
}

func (route *socialMediaRoute) Delete(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")

	if err := route.socialMediaUseCase.Delete(ctx.Request.Context(), socialMediaID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}
