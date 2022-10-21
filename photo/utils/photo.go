package utils

import (
	"time"
)

type Photo struct {
	ID        string     `json:"id"`
	Title     string     `json:"title,"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      *User      `json:"user"`
}

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type NewPhoto struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type UpdatedPhoto struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    string     `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}
