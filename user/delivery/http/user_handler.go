package delivery

import (
	"mygram-api/domain"
	"mygram-api/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type userRoute struct {
	userUseCase domain.UserUseCase
}

func NewUserRoute(handlers *gin.Engine, userUseCase domain.UserUseCase) {
	route := &userRoute{userUseCase}

	handler := handlers.Group("/users")
	{
		handler.POST("/register", route.Register)
		handler.POST("/login", route.Login)
		handler.PUT("/", route.Update)
		handler.DELETE("/", route.Delete)
	}
}

func (route *userRoute) Register(ctx *gin.Context) {
	var (
		user domain.User
		err  error
	)

	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	if err = route.userUseCase.Register(ctx.Request.Context(), &user); err != nil {
		if strings.Contains(err.Error(), "idx_users_username") {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "The username you entered has been used",
			})

			return
		}

		if strings.Contains(err.Error(), "idx_users_email") {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "The Email you entered has been used",
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
		"username": user.Username,
		"age":      user.Age,
	})
}

func (route *userRoute) Login(ctx *gin.Context) {
	var (
		user  domain.User
		err   error
		token string
	)

	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	if err = route.userUseCase.Login(ctx.Request.Context(), &user); err != nil {
		if strings.Contains(err.Error(), "invalid password") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
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

	if token, err = helpers.GenerateToken(user.ID, user.Email); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (route *userRoute) Update(c *gin.Context) {
	var (
		user domain.User
		err  error
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	updatedUser := domain.User{
		Username: user.Username,
		Email:    user.Email,
	}

	user, err = route.userUseCase.Update(c.Request.Context(), updatedUser, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})
}

func (route *userRoute) Delete(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := string(userData["id"].(string))

	err := route.userUseCase.Delete(c, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Your account has been successfully deleted",
		},
	)
}
