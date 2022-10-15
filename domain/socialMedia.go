package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"NOT NULL;type:varchar(255);"`
	SocialMediaURL string `json:"social_media_url" gorm:"NOT NULL;type:text;"`
	UserID         uint   `json:"user_id"`
}

type SosmedUseCase interface {
	CreateSosmedUC(ctx *gin.Context) (*SocialMedia, error)
	GetSosmedsUC(ctx *gin.Context) ([]*SocialMedia, error)
	UpdateSosmedUC(ctx *gin.Context) (*SocialMedia, error)
	DeleteSosmedUC(ctx *gin.Context) (*SocialMedia, error)
}

type SosmedRepository interface {
	CreateSosmedRepository(ctx *gin.Context) (*SocialMedia, error)
	GetSosmedsRepository(ctx *gin.Context) ([]*SocialMedia, error)
	UpdateSosmedRepository(ctx *gin.Context) (*SocialMedia, error)
	DeleteSosmedRepository(ctx *gin.Context) (*SocialMedia, error)
}
