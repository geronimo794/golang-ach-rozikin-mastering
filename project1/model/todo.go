package model

type Todo struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Priority string `json:"priority" gorm:"type:ENUM('low', 'medium', 'high'); default:'medium'"`
	Status   int    `json:"status"`
}
