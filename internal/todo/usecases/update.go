package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"todo-api/internal/todo/domain"
)

type UpdateUseCase struct {
	todos domain.ToDoRepository
}

type UpdateCommand struct {
	CreateCommand
	ID uuid.UUID `json:"id"`
}

func NewUpdateUseCase(todos domain.ToDoRepository) *UpdateUseCase {
	return &UpdateUseCase{
		todos: todos,
	}
}

func (useCase *UpdateUseCase) Handle(ctx context.Context, command *UpdateCommand) error {
	return useCase.updateToDo(ctx, command)
}

func (useCase *UpdateUseCase) updateToDo(ctx context.Context, command *UpdateCommand) error {
	todo, err := useCase.todos.FindByIDForUpdate(ctx, command.ID)
	if err != nil {
		return fmt.Errorf("todo: update todo %w", err)
	}

	todo.ID = command.ID
	todo.Name = command.Name
	todo.Body = command.Body
	todo.Priority = command.Priority
	todo.Deadline = command.Deadline
	todo.UpdatedAt = time.Now()

	if err = todo.Validate(ctx); err != nil {
		return fmt.Errorf("todo: validate todo %w", err)
	}

	if err = useCase.todos.Save(ctx, todo); err != nil {
		return fmt.Errorf("todo: save todo %w", err)
	}

	return nil
}
