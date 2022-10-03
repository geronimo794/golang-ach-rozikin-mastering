package web

type RequestUser struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}
