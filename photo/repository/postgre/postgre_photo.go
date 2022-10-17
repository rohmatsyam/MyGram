package postgre

import (
	"errors"
	"final_zoom/domain"
	"final_zoom/helpers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type photoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) domain.PhotoRepository {
	return &photoRepository{
		DB: db,
	}
}

var appJSON = "application/json"

func (m *photoRepository) CreatePhotoRepository(c *gin.Context) (photo *domain.Photo, err error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.UserID = userID

	err = m.DB.Debug().Create(&photo).Error

	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (m *photoRepository) GetPhotosRepository(c *gin.Context) (user *domain.User, err error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	err = m.DB.Preload("Photo").Where("id=?", userID).Find(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *photoRepository) UpdatePhotoRepository(c *gin.Context) (photo *domain.Photo, err error) {
	var newPhoto domain.Photo
	id := c.Param("photoId")
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&newPhoto)
	} else {
		c.ShouldBind(&newPhoto)
	}

	err = m.DB.Debug().Model(&photo).Where("id=?", id).First(&photo).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	photo.Title = newPhoto.Title
	photo.Caption = newPhoto.Caption
	photo.PhotoURL = newPhoto.PhotoURL

	err = m.DB.Debug().Model(&photo).Updates(&newPhoto).Error
	if err != nil {
		return nil, errors.New("update failed")
	}
	return photo, nil
}

func (m *photoRepository) DeletePhotoRepository(c *gin.Context) (photo *domain.Photo, err error) {
	id := c.Param("photoId")

	err = m.DB.First(&photo, "id=?", id).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	err = m.DB.Unscoped().Delete(&photo).Error
	if err != nil {
		return nil, errors.New("delete failed")
	}
	return photo, nil
}
