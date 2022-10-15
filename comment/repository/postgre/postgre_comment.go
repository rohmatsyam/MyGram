package postgre

import (
	"errors"
	"final_zoom/domain"
	"final_zoom/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) domain.CommentRepository {
	return &commentRepository{
		DB: db,
	}
}

var appJSON = "application/json"

func (m *commentRepository) CreateCommentRepository(c *gin.Context) (comment *domain.Comment, err error) {
	contentType := helpers.GetContentType(c)

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	err = m.DB.Debug().Create(&comment).Error

	if err != nil {
		return nil, err
	}
	return comment, nil
}
func (m *commentRepository) GetCommentsRepository(c *gin.Context) (comments []*domain.Comment, err error) {
	err = m.DB.Debug().Model(&comments).Find(&comments).Error

	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (m *commentRepository) UpdateCommentRepository(c *gin.Context) (comment *domain.Comment, err error) {
	var newComment domain.Comment
	id := c.Param("commentId")
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&newComment)
	} else {
		c.ShouldBind(&newComment)
	}

	err = m.DB.Debug().Model(&comment).Where("id=?", id).First(&comment).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	comment.Message = newComment.Message

	err = m.DB.Debug().Model(&comment).Updates(&newComment).Error
	if err != nil {
		return nil, errors.New("update failed")
	}
	return comment, nil
}
func (m *commentRepository) DeleteCommentRepository(c *gin.Context) (comment *domain.Comment, err error) {
	id := c.Param("commentId")

	err = m.DB.First(&comment, "id=?", id).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	err = m.DB.Delete(&comment).Error
	if err != nil {
		return nil, errors.New("delete failed")
	}
	return comment, nil
}
