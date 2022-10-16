package domain

import (
	"context"
	"time"
)

type User struct {
	ID           string        `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	Username     string        `gorm:"type:VARCHAR(50);unique;not null" valid:"required" form:"username" json:"username" binding:"required"`
	Email        string        `gorm:"type:VARCHAR(50);unique;not null" valid:"email,required" form:"email" json:"email" binding:"required"`
	Password     string        `gorm:"not null" valid:"required,minstringlength(6)" form:"password" json:"password" binding:"required"`
	Age          uint          `gorm:"not null" valid:"required" validate:"min=8" form:"age" json:"age" binding:"required"`
	CreatedAt    *time.Time    `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt    *time.Time    `gorm:"not null;autocreateTime" json:"updatedAt"`
	Photos       []Photo       `json:"photos"`
	SocialMedias []SocialMedia `json:"socialMedias"`
}

type UserUseCase interface {
	AddUserUseCase(context.Context, *User) error
	GetUsersUseCase(ctx context.Context) ([]User, error)
	DeleteUserUseCase(ctx context.Context, id string) error
}

type UserRepository interface {
	AddUser(ctx context.Context, user *User) error
	GetUsers(ctx context.Context) (users []User, err error)
	DeleteUser(ctx context.Context, id string) error
}
