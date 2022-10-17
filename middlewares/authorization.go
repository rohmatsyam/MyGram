package middlewares

import (
	"final_zoom/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func UserAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		User := domain.User{}
		err = db.Select("id").First(&User, uint(userId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if User.ID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you're not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func PhotoAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		Photo := domain.Photo{}
		err = db.Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you're not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func CommentAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		Comment := domain.Comment{}
		err = db.Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you're not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func SosmedAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		SocialMedia := domain.SocialMedia{}
		err = db.Select("user_id").First(&SocialMedia, uint(socialMediaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you're not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
