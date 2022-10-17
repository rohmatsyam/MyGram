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

type sosmedRepository struct {
	DB *gorm.DB
}

func NewSosmedRepository(db *gorm.DB) domain.SosmedRepository {
	return &sosmedRepository{
		DB: db,
	}
}

var appJSON = "application/json"

func (m *sosmedRepository) CreateSosmedRepository(c *gin.Context) (sosmed *domain.SocialMedia, err error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)

	if contentType == appJSON {
		c.ShouldBindJSON(&sosmed)
	} else {
		c.ShouldBind(&sosmed)
	}

	sosmed.UserID = userID
	err = m.DB.Debug().Create(&sosmed).Error

	if err != nil {
		return nil, err
	}
	return sosmed, nil

}

func (m *sosmedRepository) GetSosmedsRepository(c *gin.Context) (results []map[string]interface{}, err error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	query := fmt.Sprintf(`
	SELECT s.id AS id_sosmed,s.name,s.social_media_url,s.user_id,s.created_at,s.updated_at,
	u.id AS id_user,u.username,u.email
	FROM social_media s
	LEFT JOIN users u on s.user_id = u.id WHERE u.id=%d`, userID)
	err = m.DB.Debug().Raw(query).Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (m *sosmedRepository) UpdateSosmedRepository(c *gin.Context) (sosmed *domain.SocialMedia, err error) {
	var newSosmed domain.SocialMedia
	id := c.Param("socialMediaId")
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&newSosmed)
	} else {
		c.ShouldBind(&newSosmed)
	}

	err = m.DB.Debug().Model(&sosmed).Where("id=?", id).First(&sosmed).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	sosmed.Name = newSosmed.Name
	sosmed.SocialMediaURL = newSosmed.SocialMediaURL

	err = m.DB.Debug().Model(&sosmed).Updates(&newSosmed).Error
	if err != nil {
		return nil, errors.New("update failed")
	}
	return sosmed, nil
}

func (m *sosmedRepository) DeleteSosmedRepository(c *gin.Context) (sosmed *domain.SocialMedia, err error) {
	id := c.Param("socialMediaId")

	err = m.DB.First(&sosmed, "id=?", id).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	err = m.DB.Unscoped().Delete(&sosmed).Error
	if err != nil {
		return nil, errors.New("delete failed")
	}
	return sosmed, nil
}
