package usecase

import (
	"final_zoom/domain"

	"github.com/gin-gonic/gin"
)

type photoUseCase struct {
	photoRepo domain.PhotoRepository
}

func NewPhotoUseCase(repo domain.PhotoRepository) domain.PhotoUseCase {
	return photoUseCase{
		photoRepo: repo,
	}
}

func (c photoUseCase) CreatePhotoUC(ctx *gin.Context) (photo *domain.Photo, err error) {
	return c.photoRepo.CreatePhotoRepository(ctx)
}

func (c photoUseCase) GetPhotosUC(ctx *gin.Context) (photos []*domain.Photo, err error) {
	return c.photoRepo.GetPhotosRepository(ctx)
}

func (c photoUseCase) UpdatePhotoUC(ctx *gin.Context) (photo *domain.Photo, err error) {
	return c.photoRepo.UpdatePhotoRepository(ctx)
}

func (c photoUseCase) DeletePhotoUC(ctx *gin.Context) (photo *domain.Photo, err error) {
	return c.photoRepo.DeletePhotoRepository(ctx)
}
