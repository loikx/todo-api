package domain

import (
	"context"
	"fmt"
)

func (todo *ToDo) Validate(ctx context.Context) error {
	if len(todo.Name) == 0 {
		return fmt.Errorf("empty name")
	}

	if len(todo.Body) > 1000 {
		return fmt.Errorf("len of body nust be lower")
	}

	return nil
}
