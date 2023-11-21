package usecases

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"todo-api/internal/todo/domain"
)

type FindUseCase struct {
	todos domain.ToDoRepository
}

func NewFindUseCase(todos domain.ToDoRepository) *FindUseCase {
	return &FindUseCase{
		todos: todos,
	}
}

func (useCase *FindUseCase) Handle(ctx context.Context, id uuid.UUID) (*Response, error) {
	todo, err := useCase.todos.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("todo: find by id %w", err)
	}

	return &Response{Todo: todo}, nil
}
