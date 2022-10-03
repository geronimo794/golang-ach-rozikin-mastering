package web

type RequestTodo struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Priority string `json:"priority" form:"priority" validate:"required,oneof=low medium high"`
}
