package model

import "time"

type Author struct {
	Id         string    `json:"id"`
	Fullname   string    `json:"fullname"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Role       string    `json:"role"`
}
