package http

import (
	"final_zoom/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SosmedHandler struct {
	sosmedUseCase domain.SosmedUseCase
}

func NewSosmedHandler(r *gin.Engine, sosmedUc domain.SosmedUseCase) {
	handler := &SosmedHandler{
		sosmedUseCase: sosmedUc,
	}
	router := r.Group("/socialmedias")
	router.POST("/", handler.CreateSosmed)
	router.GET("/", handler.GetSosmeds)
	router.PUT("/:socialMediaId", handler.UpdateSosmed)
	router.DELETE("/:socialMediaId", handler.DeleteSosmed)
}

func (h SosmedHandler) CreateSosmed(c *gin.Context) {
	res, err := h.sosmedUseCase.CreateSosmedUC(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               res.ID,
		"name":             res.Name,
		"social_media_url": res.SocialMediaURL,
		"user_id":          res.UserID,
		"created_at":       res.CreatedAt,
	})
}

func (h SosmedHandler) GetSosmeds(c *gin.Context) {
	res, err := h.sosmedUseCase.GetSosmedsUC(c)
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

func (h SosmedHandler) UpdateSosmed(c *gin.Context) {
	res, err := h.sosmedUseCase.UpdateSosmedUC(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":               res.ID,
		"name":             res.Name,
		"social_media_url": res.SocialMediaURL,
		"user_id":          res.UserID,
		"updated_at":       res.UpdatedAt,
	})
}

func (h SosmedHandler) DeleteSosmed(c *gin.Context) {
	_, err := h.sosmedUseCase.DeleteSosmedUC(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
