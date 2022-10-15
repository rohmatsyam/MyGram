package usecase

import (
	"final_zoom/domain"

	"github.com/gin-gonic/gin"
)

type commentUseCase struct {
	commentRepo domain.CommentRepository
}

func NewCommentUseCase(repo domain.CommentRepository) domain.CommentUseCase {
	return commentUseCase{
		commentRepo: repo,
	}
}

func (c commentUseCase) CreateCommentUC(ctx *gin.Context) (comment *domain.Comment, err error) {
	return c.commentRepo.CreateCommentRepository(ctx)
}

func (c commentUseCase) GetCommentsUC(ctx *gin.Context) (comment []*domain.Comment, err error) {
	return c.commentRepo.GetCommentsRepository(ctx)
}

func (c commentUseCase) UpdateCommentUC(ctx *gin.Context) (comment *domain.Comment, err error) {
	return c.commentRepo.UpdateCommentRepository(ctx)
}

func (c commentUseCase) DeleteCommentUC(ctx *gin.Context) (comment *domain.Comment, err error) {
	return c.commentRepo.DeleteCommentRepository(ctx)
}
