package http

import (
	"final_zoom/domain"
	"final_zoom/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoHandler struct {
	photoUseCase domain.PhotoUseCase
}

func NewPhotoHandler(r *gin.Engine, photoUc domain.PhotoUseCase, db *gorm.DB) {
	handler := &PhotoHandler{
		photoUseCase: photoUc,
	}
	router := r.Group("/photos")
	{
		router.Use(middlewares.Authentication())
		router.POST("/", handler.CreatePhoto)
		router.GET("/", handler.GetPhotos)
		router.PUT("/:photoId", middlewares.PhotoAuthorization(db), handler.UpdatePhoto)
		router.DELETE("/:photoId", middlewares.PhotoAuthorization(db), handler.DeletePhoto)
	}
}

func (h PhotoHandler) CreatePhoto(c *gin.Context) {
	res, err := h.photoUseCase.CreatePhotoUC(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         res.ID,
		"title":      res.Title,
		"caption":    res.Caption,
		"photo_url":  res.PhotoURL,
		"user_id":    res.UserID,
		"created_at": res.CreatedAt,
	})
}

func (h PhotoHandler) GetPhotos(c *gin.Context) {
	res, err := h.photoUseCase.GetPhotosUC(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": err.Error(),
			"count":  0,
		})
		return
	}

	hasil := make([]map[string]interface{}, len(res.Photo))
	for i := 0; i < len(res.Photo); i++ {
		hasil[i] = map[string]interface{}{
			"id":         res.Photo[i].ID,
			"title":      res.Photo[i].Title,
			"caption":    res.Photo[i].Caption,
			"photo_url":  res.Photo[i].PhotoURL,
			"user_id":    res.Photo[i].UserID,
			"created_at": res.Photo[i].CreatedAt,
			"updated_at": res.Photo[i].UpdatedAt,
			"User": map[string]interface{}{
				"email":    res.Email,
				"username": res.Username,
			},
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"count":  len(hasil),
		"result": hasil,
	})
}

func (h PhotoHandler) UpdatePhoto(c *gin.Context) {
	res, err := h.photoUseCase.UpdatePhotoUC(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         res.ID,
		"title":      res.Title,
		"caption":    res.Caption,
		"photo_url":  res.PhotoURL,
		"user_id":    res.UserID,
		"updated_at": res.UpdatedAt,
	})
}

func (h PhotoHandler) DeletePhoto(c *gin.Context) {
	_, err := h.photoUseCase.DeletePhotoUC(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
