package domain

import (
	"time"
)

type User struct {
	Id             string    `json:"id"`
	Login          string    `json:"login"`
	Password       string    `json:"password"`
	Username       string    `json:"username"`
	DateOfCreation time.Time `json:"date_of_creation"`
}

type SignInInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
