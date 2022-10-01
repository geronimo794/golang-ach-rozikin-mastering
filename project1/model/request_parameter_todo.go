package model

type RequestParameterTodo struct {
	Keyword string `json:"keyword"`
	Status  int    `json:"status"`
}
