package http

import (
	"final_zoom/domain"
	"final_zoom/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentHandler struct {
	commentUseCase domain.CommentUseCase
}

func NewCommentHandler(r *gin.Engine, commentUc domain.CommentUseCase, db *gorm.DB) {
	handler := &CommentHandler{
		commentUseCase: commentUc,
	}
	router := r.Group("comments")
	{
		router.Use(middlewares.Authentication())
		router.POST("/", handler.CreateComment)
		router.GET("/", handler.GetComments)
		router.PUT("/:commentId", middlewares.CommentAuthorization(db), handler.UpdateComment)
		router.DELETE("/:commentId", middlewares.CommentAuthorization(db), handler.DeleteComment)
	}
}

func (h CommentHandler) CreateComment(c *gin.Context) {
	res, err := h.commentUseCase.CreateCommentUC(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         res.ID,
		"message":    res.Message,
		"photo_id":   res.PhotoID,
		"user_id":    res.UserID,
		"created_at": res.CreatedAt,
	})
}

func (h CommentHandler) GetComments(c *gin.Context) {
	res, err := h.commentUseCase.GetCommentsUC(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": err.Error(),
			"count":  0,
		})
		return
	}

	hasil := make([]map[string]interface{}, len(res.Comment))
	for i := 0; i < len(res.Comment); i++ {
		hasil[i] = map[string]interface{}{
			"id":         res.Comment[i].ID,
			"message":    res.Comment[i].Message,
			"photo_id":   res.Comment[i].PhotoID,
			"user_id":    res.Comment[i].UserID,
			"created_at": res.Comment[i].CreatedAt,
			"updated_at": res.Comment[i].UpdatedAt,
			"User": map[string]interface{}{
				"id":       res.ID,
				"email":    res.Email,
				"username": res.Username,
			},
			"Photo": map[string]interface{}{
				"id":        res.Photo[i].ID,
				"title":     res.Photo[i].Title,
				"caption":   res.Photo[i].Caption,
				"photo_url": res.Photo[i].PhotoURL,
				"user_id":   res.Photo[i].UserID,
			},
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"count":  len(hasil),
		"result": hasil,
	})
}

func (h CommentHandler) UpdateComment(c *gin.Context) {
	res, err := h.commentUseCase.UpdateCommentUC(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         res.ID,
		"message":    res.Message,
		"photo_id":   res.PhotoID,
		"user_id":    res.UserID,
		"updated_at": res.UpdatedAt,
	})
}

func (h CommentHandler) DeleteComment(c *gin.Context) {
	_, err := h.commentUseCase.DeleteCommentUC(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
