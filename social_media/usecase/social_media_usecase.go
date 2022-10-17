package usecase

import (
	"final_zoom/domain"

	"github.com/gin-gonic/gin"
)

type sosmedUseCase struct {
	sosmedRepo domain.SosmedRepository
}

func NewSosmedUseCase(repo domain.SosmedRepository) domain.SosmedUseCase {
	return sosmedUseCase{
		sosmedRepo: repo,
	}
}

func (c sosmedUseCase) CreateSosmedUC(ctx *gin.Context) (sosmed *domain.SocialMedia, err error) {
	return c.sosmedRepo.CreateSosmedRepository(ctx)
}

func (c sosmedUseCase) GetSosmedsUC(ctx *gin.Context) (user *domain.User, err error) {
	return c.sosmedRepo.GetSosmedsRepository(ctx)
}

func (c sosmedUseCase) UpdateSosmedUC(ctx *gin.Context) (sosmed *domain.SocialMedia, err error) {
	return c.sosmedRepo.UpdateSosmedRepository(ctx)
}

func (c sosmedUseCase) DeleteSosmedUC(ctx *gin.Context) (sosmed *domain.SocialMedia, err error) {
	return c.sosmedRepo.DeleteSosmedRepository(ctx)
}
