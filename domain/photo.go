package domain

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string    `json:"title" gorm:"NOT NULL;type:varchar(255);" valid:"required"`
	Caption  string    `json:"caption" gorm:"type:varchar(255);"`
	PhotoURL string    `json:"photo_url" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID   uint      `json:"user_id"`
	Comment  []Comment `json:"comments,omitempty" gorm:"foreignKey:photo_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
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
	GetPhotosUC(ctx *gin.Context) ([]*Photo, error)
	UpdatePhotoUC(ctx *gin.Context) (*Photo, error)
	DeletePhotoUC(ctx *gin.Context) (*Photo, error)
}

type PhotoRepository interface {
	CreatePhotoRepository(ctx *gin.Context) (*Photo, error)
	GetPhotosRepository(ctx *gin.Context) ([]*Photo, error)
	UpdatePhotoRepository(ctx *gin.Context) (*Photo, error)
	DeletePhotoRepository(ctx *gin.Context) (*Photo, error)
}
