package utils

import (
	"time"
)

type ResponseData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Photo struct {
	ID        string     `json:"id"`
	Title     string     `json:"title,"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photoUrl"`
	UserID    string     `json:"userId"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
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
	PhotoUrl  string     `json:"photoUrl"`
	UserID    string     `json:"userId"`
	CreatedAt *time.Time `json:"createdAt"`
}

type UpdatedPhoto struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photoUrl"`
	UserID    string     `json:"userId"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
