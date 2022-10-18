package domain

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"not null" json:"userId"`
	User      User       `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	PhotoID   uint       `gorm:"not null" form:"photoId" json:"photoId"`
	Photo     Photo      `gorm:"foreignKey:PhotoID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	Message   string     `gorm:"not null" valid:"required" form:"message" json:"message"`
	CreatedAt *time.Time `gorm:"not null;autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"not null;autoCreateTime" json:"updatedAt,omitempty"`
}

func (c *Comment) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}

	return
}

func (c *Comment) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}
	return
}

type CommentUseCase interface {
	Fetch(context.Context, *[]Comment, uint) error
	Store(context.Context, *Comment) error
	GetByUserID(context.Context, *Comment, uint) error
	Update(context.Context, Comment, uint) (Comment, error)
	Delete(context.Context, uint) error
}

type CommentRepository interface {
	Fetch(context.Context, *[]Comment, uint) error
	Store(context.Context, *Comment) error
	GetByUserID(context.Context, *Comment, uint) error
	Update(context.Context, Comment, uint) (Comment, error)
	Delete(context.Context, uint) error
}
