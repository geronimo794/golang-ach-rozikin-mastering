package web

type RequestParameterTodo struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  int8   `json:"status" form:"status" validate:"oneof=0 1"`
}
