package model

type Todo struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Priority string `json:"priority" gorm:"type:ENUM('low', 'medium', 'high'); default:'medium'" validate:"required,oneof=low medium high"`
	IsDone   bool   `json:"is_done"`
}
