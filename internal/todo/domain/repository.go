package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

type ToDoRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*ToDo, error)
	FindByIDForUpdate(ctx context.Context, id uuid.UUID) (*ToDo, error)
	Save(ctx context.Context, todo *ToDo) error
}
