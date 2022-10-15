package usecase

import (
	"errors"
	"final_zoom/domain"
	"regexp"

	"github.com/gin-gonic/gin"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return userUseCase{
		userRepo: repo,
	}
}

func (c userUseCase) UserRegisterUc(ctx *gin.Context) (user *domain.User, err error) {
	return c.userRepo.UserRegisterRepository(ctx)
}

func (c userUseCase) UserLoginUc(ctx *gin.Context) (token string, err error) {
	return c.userRepo.UserLoginRepository(ctx)
}

func (c userUseCase) GetUserByIdUc(ctx *gin.Context) (user *domain.User, err error) {
	err = CheckParamIsNumber(ctx.Param("userId"))
	if err != nil {
		return nil, errors.New("user not found")
	}
	return c.userRepo.GetUserByIdRepository(ctx)
}

func (c userUseCase) GetUsersUc(ctx *gin.Context) ([]*domain.User, error) {
	return c.userRepo.GetUsersRepository(ctx)
}

func (c userUseCase) UpdateUserUc(ctx *gin.Context) (*domain.User, error) {
	return c.userRepo.UpdateUserRepository(ctx)
}

func (c userUseCase) DeleteUserUc(ctx *gin.Context) (*domain.User, error) {
	return c.userRepo.DeleteUserRepository(ctx)
}

func CheckParamIsNumber(param string) error {
	sampleRegexp := regexp.MustCompile(`\d`)
	match := sampleRegexp.MatchString(param)
	if !match {
		return errors.New("not a number")
	}
	return nil
}
