package domain

import "context"

type SocialMedia struct {
	ID             string `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	Name           string `gorm:"type:VARCHAR(50);not null" valid:"required" form:"name" json:"name"`
	SocialMediaURL string `gorm:"not null" valid:"required" form:"socialMediaUrl" json:"socialMediaUrl"`
	UserID         string `gorm:"type:VARCHAR(50);not null" json:"userId"`
	User           User   `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt      string `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt      string `gorm:"type:timestamp;not null" json:"updatedAt"`
}

type SocialMediaUseCase interface {
	AddSocialMediaUseCase(context.Context, *SocialMedia) error
	GetSocialMediasUseCase(ctx context.Context) ([]SocialMedia, error)
	UpdateSocialMediaUseCase(ctx context.Context, socialMedia *SocialMedia) error
	DeleteSocialMediaUseCase(ctx context.Context, id string) error
}

type SocialMediaRepository interface {
	AddSocialMedia(ctx context.Context, socialMedia *SocialMedia) error
	GetSocialMedias(ctx context.Context) (socialMedia []SocialMedia, err error)
	UpdateSocialMedia(ctx context.Context, socialMedia *SocialMedia) error
	DeleteSocialMedia(ctx context.Context, id string) error
}
