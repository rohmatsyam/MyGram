package domain

import (
	"final_zoom/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string        `json:"username" gorm:"NOT NULL;unique;type:varchar(255);" valid:"required"`
	Email       string        `json:"email" gorm:"NOT NULL;unique;type:varchar(255);" valid:"required,email"`
	Password    string        `json:"password" gorm:"NOT NULL;type:text;" valid:"required,minstringlength(6)"`
	Age         uint          `json:"age" gorm:"NOT NULL;type:integer;" valid:"required,range(8|100)"`
	SocialMedia []SocialMedia `json:"social_media,omitempty" gorm:"foreignKey:user_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photo       []Photo       `json:"photos,omitempty" gorm:"foreignKey:user_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comment     []Comment     `json:"comments,omitempty" gorm:"foreignKey:user_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}

type UserUseCase interface {
	UserRegisterUc(ctx *gin.Context) (*User, error)
	UserLoginUc(ctx *gin.Context) (string, error)
	GetUserByIdUc(ctx *gin.Context) (*User, error)
	GetUsersUc(ctx *gin.Context) ([]*User, error)
	UpdateUserUc(ctx *gin.Context) (*User, error)
	DeleteUserUc(ctx *gin.Context) (*User, error)
}

type UserRepository interface {
	UserRegisterRepository(ctx *gin.Context) (*User, error)
	UserLoginRepository(ctx *gin.Context) (string, error)
	GetUserByIdRepository(ctx *gin.Context) (*User, error)
	GetUsersRepository(ctx *gin.Context) ([]*User, error)
	UpdateUserRepository(ctx *gin.Context) (*User, error)
	DeleteUserRepository(ctx *gin.Context) (*User, error)
}
