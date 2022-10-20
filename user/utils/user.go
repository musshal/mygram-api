package utils

import "time"

type UpdatedUser struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Age       uint       `json:"age"`
	UpdatedAt *time.Time `json:"updatedAt"`
}