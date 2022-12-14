package http

import (
	"final_zoom/domain"
	"final_zoom/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(r *gin.Engine, userUc domain.UserUseCase, db *gorm.DB) {
	handler := &UserHandler{
		userUseCase: userUc,
	}
	router := r.Group("/users")
	router.POST("/register", handler.RegisterUser)
	router.POST("/login", handler.LoginUser)
	{
		router.Use(middlewares.Authentication())
		router.GET("/:userId", handler.GetUserById)
		router.GET("/", handler.GetUsers)
		router.PUT("/:userId", middlewares.UserAuthorization(db), handler.UpdateUser)
		router.DELETE("/:userId", middlewares.UserAuthorization(db), handler.DeleteUser)
	}
}
func (h UserHandler) GetUserById(c *gin.Context) {
	var result gin.H
	res, err := h.userUseCase.GetUserByIdUc(c)
	if err != nil {
		result = gin.H{
			"count":  0,
			"result": err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	// hasil := make([]map[string]interface{}, len(res.SocialMedia))
	// for i := 0; i < len(res.SocialMedia); i++ {
	// 	hasil[i] = map[string]interface{}{
	// 		"id":               res.SocialMedia[i].ID,
	// 		"name":             res.SocialMedia[i].Name,
	// 		"social_media_url": res.SocialMedia[i].SocialMediaURL,
	// 		"created_at":       res.SocialMedia[i].CreatedAt,
	// 		"updated_at":       res.SocialMedia[i].UpdatedAt,
	// 		"User": map[string]interface{}{
	// 			"id":       res.ID,
	// 			"email":    res.Email,
	// 			"username": res.Username,
	// 		},
	// 	}
	// }
	res.Password = ""
	c.JSON(http.StatusFound, gin.H{
		"count":  1,
		"result": res,
	})
}

func (h UserHandler) GetUsers(c *gin.Context) {
	res, err := h.userUseCase.GetUsersUc(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": err.Error(),
			"count":  0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":  len(res),
		"result": res,
	})
}

func (h UserHandler) RegisterUser(c *gin.Context) {
	res, err := h.userUseCase.UserRegisterUc(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       res.ID,
		"email":    res.Email,
		"username": res.Username,
		"age":      res.Age,
	})
}

func (h UserHandler) LoginUser(c *gin.Context) {
	token, err := h.userUseCase.UserLoginUc(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h UserHandler) UpdateUser(c *gin.Context) {
	res, err := h.userUseCase.UpdateUserUc(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         res.ID,
		"email":      res.Email,
		"username":   res.Username,
		"age":        res.Age,
		"updated_at": res.UpdatedAt,
	})
}

func (h UserHandler) DeleteUser(c *gin.Context) {
	_, err := h.userUseCase.DeleteUserUc(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
