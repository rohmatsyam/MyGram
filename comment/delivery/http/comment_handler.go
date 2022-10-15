package http

import (
	"final_zoom/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentUseCase domain.CommentUseCase
}

func NewCommentHandler(r *gin.Engine, commentUc domain.CommentUseCase) {
	handler := &CommentHandler{
		commentUseCase: commentUc,
	}
	router := r.Group("comments")
	router.POST("/", handler.CreateComment)
	router.GET("/", handler.GetComments)
	router.PUT("/:commentId", handler.UpdateComment)
	router.DELETE("/:commentId", handler.DeleteComment)
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

	c.JSON(http.StatusOK, gin.H{
		"result": res,
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
