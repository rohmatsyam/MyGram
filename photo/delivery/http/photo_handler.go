package http

import (
	"final_zoom/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoHandler struct {
	photoUseCase domain.PhotoUseCase
}

func NewPhotoHandler(r *gin.Engine, photoUc domain.PhotoUseCase) {
	handler := &PhotoHandler{
		photoUseCase: photoUc,
	}
	router := r.Group("/photos")
	router.POST("/", handler.CreatePhoto)
	router.GET("/", handler.GetPhotos)
	router.PUT("/:photoId", handler.UpdatePhoto)
	router.DELETE("/:photoId", handler.DeletePhoto)
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

	if len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data masih kosong",
			"count":   0,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": res,
		"count":  len(res),
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
