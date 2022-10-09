package model

import (
	"fmt"
	"io"
	"strconv"
)

type Todo struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Priority string `json:"priority" gorm:"type:ENUM('low', 'medium', 'high'); default:'medium'" validate:"required,oneof=low medium high"`
	IsDone   bool   `json:"is_done"`
}

type TodoInput struct {
	Name     string       `json:"name"`
	Priority TodoPriority `json:"priority"`
}

type TodoPriority string

const (
	TodoPriorityLow    TodoPriority = "low"
	TodoPriorityMedium TodoPriority = "medium"
	TodoPriorityHigh   TodoPriority = "high"
)

var AllTodoPriority = []TodoPriority{
	TodoPriorityLow,
	TodoPriorityMedium,
	TodoPriorityHigh,
}

func (e TodoPriority) IsValid() bool {
	switch e {
	case TodoPriorityLow, TodoPriorityMedium, TodoPriorityHigh:
		return true
	}
	return false
}

func (e TodoPriority) String() string {
	return string(e)
}

func (e *TodoPriority) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoPriority(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoPriority", str)
	}
	return nil
}

func (e TodoPriority) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
