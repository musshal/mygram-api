package utils

import "time"

type SocialMedias struct {
	SocialMedias interface{} `json:"social_medias"`
}

type NewSocialMedia struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         string     `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type UpdatedSocialMedia struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         string     `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
