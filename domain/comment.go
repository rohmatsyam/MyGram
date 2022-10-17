package domain

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message string `json:"message" gorm:"NOT NULL;type:text;" valid:"required"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}

type CommentUseCase interface {
	CreateCommentUC(ctx *gin.Context) (*Comment, error)
	GetCommentsUC(ctx *gin.Context) (*User, error)
	UpdateCommentUC(ctx *gin.Context) (*Comment, error)
	DeleteCommentUC(ctx *gin.Context) (*Comment, error)
}

type CommentRepository interface {
	CreateCommentRepository(ctx *gin.Context) (*Comment, error)
	GetCommentsRepository(ctx *gin.Context) (*User, error)
	UpdateCommentRepository(ctx *gin.Context) (*Comment, error)
	DeleteCommentRepository(ctx *gin.Context) (*Comment, error)
}
