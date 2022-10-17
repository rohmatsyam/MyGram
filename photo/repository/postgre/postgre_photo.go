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

func (m *photoRepository) GetPhotosRepository(c *gin.Context) (results []map[string]interface{}, err error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	query := fmt.Sprintf(`
	SELECT p.id AS id_photo,p.title,p.caption,p.photo_url,p.user_id,p.created_at,p.updated_at,
	u.email,u.username
	FROM photos p
	lEft JOIN users u ON p.user_id = u.id WHERE u.id=%d`, userID)
	err = m.DB.Debug().Raw(query).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
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
