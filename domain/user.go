package domain

import (
	"context"
	"mygram-api/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `gorm:"type:VARCHAR(50);uniqueIndex;not null" valid:"required" form:"username" json:"username" example:"johndoe"`
	Email        string         `gorm:"type:VARCHAR(50);uniqueIndex;not null" valid:"email,required" form:"email" json:"email" example:"johndoe@example.com"`
	Password     string         `gorm:"not null" valid:"required,minstringlength(6)" form:"password" json:"password,omitempty" example:"secret"`
	Age          uint           `gorm:"not null" valid:"required,range(8|63)" form:"age" json:"age,omitempty" example:"8"`
	CreatedAt    *time.Time     `gorm:"not null;autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt    *time.Time     `gorm:"not null;autocreateTime" json:"updatedAt,omitempty"`
	Photos       *[]Photo       `json:"-"`
	SocialMedias *[]SocialMedia `json:"-"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return err
	}

	user.Password = helpers.Hash(user.Password)

	return
}

func (user *User) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return err
	}

	return
}

type UserUseCase interface {
	Register(context.Context, *User) error
	Login(context.Context, *User) error
	Update(context.Context, User, uint) (User, error)
	Delete(context.Context, uint) error
}

type UserRepository interface {
	Register(context.Context, *User) error
	Login(context.Context, *User) error
	Update(context.Context, User, uint) (User, error)
	Delete(context.Context, uint) error
}
