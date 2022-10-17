package postgre

import (
	"errors"
	"final_zoom/domain"
	"final_zoom/helpers"

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

func (m *sosmedRepository) GetSosmedsRepository(c *gin.Context) (user *domain.User, err error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	err = m.DB.Preload("SocialMedia").Where("id=?", userID).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
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
