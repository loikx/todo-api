package usecases

import (
	"context"
	"fmt"
	"time"

	"todo-api/internal/todo/domain"
)

type CreateUseCase struct {
	todos domain.ToDoRepository
}

type CreateCommand struct {
	Priority domain.Priority `json:"priority"`
	Name     string          `json:"name"`
	Body     string          `json:"body"`
	Deadline *time.Time      `json:"deadline"`
}

func NewCreateUseCase(todos domain.ToDoRepository) *CreateUseCase {
	return &CreateUseCase{
		todos: todos,
	}
}

func (useCase *CreateUseCase) Handle(ctx context.Context, createCommand *CreateCommand) (*Response, error) {
	return useCase.createToDo(ctx, createCommand)
}

func (useCase *CreateUseCase) createToDo(ctx context.Context, command *CreateCommand) (*Response, error) {
	todo := domain.NewToDo()
	todo.Name = command.Name
	todo.Body = command.Body
	todo.Priority = command.Priority
	todo.Deadline = command.Deadline

	if err := todo.Validate(ctx); err != nil {
		return nil, fmt.Errorf("todo: validate todo %w", err)
	}

	if err := useCase.todos.Save(ctx, todo); err != nil {
		return nil, fmt.Errorf("todo: save todo %w", err)
	}

	return &Response{Todo: todo}, nil
}
