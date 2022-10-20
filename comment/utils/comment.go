package utils

import "time"

type Comment struct {
	ID        string     `json:"id"`
	Message   string     `json:"message"`
	PhotoID   string     `json:"photoId"`
	UserID    string     `json:"userId"`
	UpdatedAt *time.Time `json:"updatedAt"`
	CreatedAt *time.Time `json:"createdAt"`
	User      User       `json:"user"`
	Photo     Photo      `json:"photo"`
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Photo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl"`
	UserID   string `json:"userId"`
}

type NewComment struct {
	ID        string     `json:"id"`
	UserID    string     `json:"userId"`
	PhotoID   string     `json:"photoId"`
	Message   string     `json:"message"`
	CreatedAt *time.Time `json:"createdAt"`
}

type UpdatedComment struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photoUrl"`
	UserID    string     `json:"userId"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
