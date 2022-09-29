package model

type RequestParameterTodo struct {
	Keyword string `json:"name"`
	Status  int    `json:"status"`
}
