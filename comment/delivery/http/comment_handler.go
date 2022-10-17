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

	if len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data masih kosong",
			"count":   0,
		})
		return
	}

	hasil := make([]map[string]interface{}, len(res))
	for i := 0; i < len(res); i++ {
		hasil[i] = map[string]interface{}{
			"id":         res[i]["id_comment"],
			"message":    res[i]["message"],
			"photo_id":   res[i]["c_photo_id"],
			"user_id":    res[i]["c_user_id"],
			"created_at": res[i]["created_at"],
			"updated_at": res[i]["updated_at"],
			"User": map[string]interface{}{
				"id":       res[i]["id_user"],
				"email":    res[i]["email"],
				"username": res[i]["username"],
			},
			"Photo": map[string]interface{}{
				"id":        res[i]["id_photo"],
				"title":     res[i]["title"],
				"caption":   res[i]["caption"],
				"photo_url": res[i]["photo_url"],
				"user_id":   res[i]["p_user_id"],
			},
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": hasil,
		"count":  len(res),
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
