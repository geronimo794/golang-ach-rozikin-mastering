package model

type RequestParameterTodo struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  uint8  `json:"status" form:"status" validate:"oneof=0 1"`
}
