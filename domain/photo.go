package domain

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"NOT NULL;type:varchar(255);" valid:"required"`
	Caption   string    `json:"caption" gorm:"type:varchar(255);"`
	PhotoURL  string    `json:"photo_url" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID    uint      `json:"user_id"`
	Comment   []Comment `json:"comments,omitempty" gorm:"foreignKey:photo_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	println(govalidator.IsURL(p.PhotoURL))
	isURL := govalidator.IsURL(p.PhotoURL)
	if !isURL {
		return errors.New("url not valid")
	}

	return
}
func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}

type PhotoUseCase interface {
	CreatePhotoUC(ctx *gin.Context) (*Photo, error)
	GetPhotosUC(ctx *gin.Context) (*User, error)
	UpdatePhotoUC(ctx *gin.Context) (*Photo, error)
	DeletePhotoUC(ctx *gin.Context) (*Photo, error)
}

type PhotoRepository interface {
	CreatePhotoRepository(ctx *gin.Context) (*Photo, error)
	GetPhotosRepository(ctx *gin.Context) (*User, error)
	UpdatePhotoRepository(ctx *gin.Context) (*Photo, error)
	DeletePhotoRepository(ctx *gin.Context) (*Photo, error)
}
