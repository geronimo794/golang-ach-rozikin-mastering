package model

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
