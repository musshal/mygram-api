package domain

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"type:VARCHAR(50);not null" valid:"required" form:"name" json:"name" example:"Facebook"`
	SocialMediaURL string     `gorm:"not null" valid:"required" form:"socialMediaUrl" json:"socialMediaUrl" example:"https://www.example.com/social-media"`
	UserID         uint       `gorm:"not null" json:"userId"`
	User           User       `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt      *time.Time `gorm:"not null;autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `gorm:"not null;autoCreateTime" json:"updatedAt,omitempty"`
}

func (s *SocialMedia) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(s); err != nil {
		return err
	}

	return
}

func (s *SocialMedia) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(s); err != nil {
		return err
	}
	return
}

type SocialMediaUseCase interface {
	Fetch(context.Context, *[]SocialMedia, uint) error
	Store(context.Context, *SocialMedia) error
	GetByUserID(context.Context, *SocialMedia, uint) error
	Update(context.Context, SocialMedia, uint) (SocialMedia, error)
	Delete(context.Context, uint) error
}

type SocialMediaRepository interface {
	Fetch(context.Context, *[]SocialMedia, uint) error
	Store(context.Context, *SocialMedia) error
	GetByUserID(context.Context, *SocialMedia, uint) error
	Update(context.Context, SocialMedia, uint) (SocialMedia, error)
	Delete(context.Context, uint) error
}
