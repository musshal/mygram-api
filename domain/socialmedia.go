package domain

import (
	"context"
	"time"
)

type SocialMedia struct {
	ID             string     `gorm:"primaryKey;type:VARCHAR(50)" json:"id" binding:"required"`
	Name           string     `gorm:"type:VARCHAR(50);not null" valid:"required" form:"name" json:"name" binding:"required"`
	SocialMediaURL string     `gorm:"not null" valid:"required" form:"socialMediaUrl" json:"socialMediaUrl" binding:"required"`
	UserID         uint       `gorm:"not null" json:"userId"`
	User           User       `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt      *time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt      *time.Time `gorm:"not null;autoCreateTime" json:"updatedAt"`
}

type SocialMediaUseCase interface {
	AddSocialMedia(context.Context, *SocialMedia) error
	GetSocialMedias(ctx context.Context) ([]SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, socialMedia *SocialMedia) error
	DeleteSocialMedia(ctx context.Context, id string) error
}

type SocialMediaRepository interface {
	AddSocialMedia(ctx context.Context, socialMedia *SocialMedia) error
	GetSocialMedias(ctx context.Context) (socialMedia []SocialMedia, err error)
	UpdateSocialMedia(ctx context.Context, socialMedia *SocialMedia) error
	DeleteSocialMedia(ctx context.Context, id string) error
}
