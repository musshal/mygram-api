package domain

import "context"

type User struct {
	ID           string        `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	Username     string        `gorm:"type:VARCHAR(50);unique;not null" valid:"required" form:"username" json:"username"`
	Email        string        `gorm:"type:VARCHAR(50);unique;not null" valid:"email,required" form:"email" json:"email"`
	Password     string        `gorm:"not null" valid:"required,minstringlength(6)" form:"password" json:"password"`
	Age          uint          `gorm:"not null" valid:"required" validate:"min=8" form:"age" json:"age"`
	CreatedAt    string        `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt    string        `gorm:"type:timestamp;not null" json:"updatedAt"`
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
