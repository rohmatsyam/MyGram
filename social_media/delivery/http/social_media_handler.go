package http

import (
	"final_zoom/domain"
	"final_zoom/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SosmedHandler struct {
	sosmedUseCase domain.SosmedUseCase
}

func NewSosmedHandler(r *gin.Engine, sosmedUc domain.SosmedUseCase, db *gorm.DB) {
	handler := &SosmedHandler{
		sosmedUseCase: sosmedUc,
	}
	router := r.Group("/socialmedias")
	{
		router.Use(middlewares.Authentication())
		router.POST("/", handler.CreateSosmed)
		router.GET("/", handler.GetSosmeds)
		router.PUT("/:socialMediaId", middlewares.SosmedAuthorization(db), handler.UpdateSosmed)
		router.DELETE("/:socialMediaId", middlewares.SosmedAuthorization(db), handler.DeleteSosmed)
	}
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

	hasil := make([]map[string]interface{}, len(res.SocialMedia))
	for i := 0; i < len(res.SocialMedia); i++ {
		hasil[i] = map[string]interface{}{
			"id":               res.SocialMedia[i].ID,
			"name":             res.SocialMedia[i].Name,
			"social_media_url": res.SocialMedia[i].SocialMediaURL,
			"created_at":       res.SocialMedia[i].CreatedAt,
			"updated_at":       res.SocialMedia[i].UpdatedAt,
			"User": map[string]interface{}{
				"id":       res.ID,
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
	// KURANG ID
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
