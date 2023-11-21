package usecases

import (
	"github.com/gofrs/uuid"
	"todo-api/internal/todo/domain"
)

type Response struct {
	Todo *domain.ToDo `json:"todo"`
}

type FindByIDsResponse struct {
	Items map[uuid.UUID]*domain.ToDo `json:"items"`
}
