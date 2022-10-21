package utils

import "time"

type ResponseData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type NewUser struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	ID       string `json:"id"`
	Username string `json:"username"`
}

type NewToken struct {
	Token string `json:"token"`
}

type UpdatedUser struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Age       uint       `json:"age"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
