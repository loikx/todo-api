package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type ToDo struct {
	ID       uuid.UUID  `json:"id"`
	Priority Priority   `json:"priority"`
	Name     string     `json:"name"`
	Body     string     `json:"body"`
	Deadline *time.Time `json:"deadline"`
	isNew    bool

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewToDo() *ToDo {
	now := time.Now()

	return &ToDo{
		ID:       uuid.Must(uuid.NewV7()),
		Priority: NoPriority,
		Name:     "",
		Body:     "",
		Deadline: nil,
		isNew:    true,

		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (todo *ToDo) IsNew() bool { return todo.isNew }
