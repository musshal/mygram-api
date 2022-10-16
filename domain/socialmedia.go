package domain

type SocialMedias struct {
	ID             string `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	Name           string `gorm:"type:VARCHAR(50);not null" json:"name"`
	SocialMediaURL string `gorm:"not null" json:"socialMediaUrl"`
	UserID         string `gorm:"type:VARCHAR(50);not null" json:"userId"`
	User           User   `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt      string `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt      string `gorm:"type:timestamp;not null" json:"updatedAt"`
}

type SocialMediaRepository interface {
	AddSocialMedia()
	GetSocialMedias()
	UpdateSocialMedia()
	DeleteSocialMedia()
}
