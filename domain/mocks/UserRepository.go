package mocks

// import (
// 	"final_zoom/domain"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/mock"
// )

// type UserRepository struct {
// 	mock.Mock
// }

// func (_m *UserRepository) GetUserByIdRepository(c *gin.Context) (user domain.User, err error) {
// 	id := c.Param("userId")
// 	_m.Called()
// 	err = _m.DB.Model(&user).Where("id=?", id).First(&user).Error
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (_m *UserRepository) GetUsersRepository(c *gin.Context) (users []domain.User, err error) {
// 	err = m.DB.Model(&users).Find(&users).Error
// 	if err != nil {
// 		return users, err
// 	}

// 	return users, nil
// }
