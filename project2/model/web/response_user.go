package web

type ResponseUser struct {
	Id    int    `json:"id"`
	Email string `json:"email" form:"email" validate:"required,email"`
}
