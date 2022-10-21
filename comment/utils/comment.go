package utils

import "time"

type NewComment struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	PhotoID   string     `json:"photo_id"`
	Message   string     `json:"message"`
	CreatedAt *time.Time `json:"created_at"`
}

type UpdatedComment struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    string     `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}
