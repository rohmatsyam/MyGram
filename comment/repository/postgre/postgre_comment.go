package postgre

import (
	"errors"
	"final_zoom/domain"
	"final_zoom/helpers"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)

	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.UserID = userID
	err = m.DB.Debug().Create(&comment).Error

	if err != nil {
		return nil, err
	}
	return comment, nil
}
func (m *commentRepository) GetCommentsRepository(c *gin.Context) (results []map[string]interface{}, err error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	query := fmt.Sprintf(`
	SELECT c.id AS id_comment,c.message,c.user_id AS c_user_id,c.photo_id AS c_photo_id,c.created_at,c.updated_at,
	u.id AS id_user,u.username,u.email,
	p.id AS id_photo,p.title,p.caption,p.photo_url,p.user_id AS p_user_id
	FROM comments c
	LEFT JOIN users u on c.user_id = u.id
	LEFT Join photos p on p.user_id = u.id WHERE u.id = %d`, userID)

	err = m.DB.Debug().Raw(query).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
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

	err = m.DB.Unscoped().Delete(&comment).Error
	if err != nil {
		return nil, errors.New("delete failed")
	}
	return comment, nil
}
