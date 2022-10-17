package postgre

import (
	"errors"
	"final_zoom/domain"
	"final_zoom/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		DB: db,
	}
}

var appJSON = "application/json"

func (m *userRepository) UserRegisterRepository(c *gin.Context) (user *domain.User, err error) {
	contentType := helpers.GetContentType(c)

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err = m.DB.Debug().Create(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *userRepository) UserLoginRepository(c *gin.Context) (token string, err error) {
	contentType := helpers.GetContentType(c)
	User := domain.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	password = User.Password
	err = m.DB.Debug().Where("email=?", &User.Email).Take(&User).Error
	if err != nil {
		return "", errors.New("record not found")
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		return "", errors.New("comparing password invalid")
	}
	token = helpers.GenerateToken(User.ID, User.Email)
	return token, nil
}

func (m *userRepository) GetUserByIdRepository(c *gin.Context) (user *domain.User, err error) {
	id := c.Param("userId")
	err = m.DB.Model(&user).Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *userRepository) GetUsersRepository(c *gin.Context) (users []*domain.User, err error) {
	err = m.DB.Model(&users).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m *userRepository) UpdateUserRepository(c *gin.Context) (user *domain.User, err error) {
	var newUser domain.User
	id := c.Param("userId")
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&newUser)
	} else {
		c.ShouldBind(&newUser)
	}

	err = m.DB.Debug().Model(&user).Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	user.Username = newUser.Username
	user.Email = newUser.Email

	err = m.DB.Debug().Model(&user).Updates(&newUser).Error
	if err != nil {
		return nil, errors.New("update failed")
	}
	return user, nil
}

func (m *userRepository) DeleteUserRepository(c *gin.Context) (user *domain.User, err error) {
	id := c.Param("userId")

	err = m.DB.First(&user, "id=?", id).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	err = m.DB.Unscoped().Delete(&user).Error
	if err != nil {
		return nil, errors.New("delete failed")
	}

	return user, nil
}
