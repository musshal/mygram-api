package delivery

import (
	"mygram-api/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type userRoute struct {
	userUseCase domain.UserUseCase
}

func NewUserRoute(handlers *gin.Engine, userUseCase domain.UserUseCase) {
	route := &userRoute{userUseCase}

	handler := handlers.Group("/users")
	{
		handler.POST("/register", route.Register)
		// handler.POST("/login", route.UserLogin)
		// handler.PUT("/", middleware.Authentication(), route.UpdateUser)
		// handler.DELETE("/", middleware.Authentication(), route.DeleteUser)
	}
}

func (route *userRoute) Register(ctx *gin.Context) {
	var (
		user domain.User
		err  error
	)

	err = ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	err = route.userUseCase.Register(ctx.Request.Context(), &user)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": err.Error(),
			})

			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Email,
		"age":      user.Age,
	})
}

func (route *userRoute) Login(ctx *gin.Context) {
	return
}

func (route *userRoute) Update(ctx *gin.Context) {
	return
}

func (route *userRoute) Delete(ctx *gin.Context) {
	return
}
