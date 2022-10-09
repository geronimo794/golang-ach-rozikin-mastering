package web

type RequestParameterTodo struct {
	Keyword string `json:"keyword" form:"keyword"`
	IsDone  string `json:"is_done" form:"is_done" validate:"oneof=true false"`
}
