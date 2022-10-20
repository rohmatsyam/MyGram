package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type SocialMedia struct {
	ID             uint      `json:"id" gorm:"primarykey"`
	Name           string    `json:"name" gorm:"NOT NULL;type:varchar(255);"`
	SocialMediaURL string    `json:"social_media_url" gorm:"NOT NULL;type:text;"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SosmedUseCase interface {
	CreateSosmedUC(ctx *gin.Context) (*SocialMedia, error)
	GetSosmedsUC(ctx *gin.Context) (*User, error)
	UpdateSosmedUC(ctx *gin.Context) (*SocialMedia, error)
	DeleteSosmedUC(ctx *gin.Context) (*SocialMedia, error)
}

type SosmedRepository interface {
	CreateSosmedRepository(ctx *gin.Context) (*SocialMedia, error)
	GetSosmedsRepository(ctx *gin.Context) (*User, error)
	UpdateSosmedRepository(ctx *gin.Context) (*SocialMedia, error)
	DeleteSosmedRepository(ctx *gin.Context) (*SocialMedia, error)
}
