package repository

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"todo-api/internal/todo/domain"
)

type Repository struct {
	connection *pgx.Conn
}

func NewRepository(connection *pgx.Conn) *Repository {
	return &Repository{connection: connection}
}

func (r *Repository) FindByID(ctx context.Context, id uuid.UUID) (*domain.ToDo, error) {
	var todo domain.ToDo
	err := r.connection.QueryRow(
		ctx,
		"select id, name, body, deadline, createdAt, updatedAt from todo.todo where id=$1",
		id).Scan(&todo.ID, &todo.Name, &todo.Body, &todo.Deadline, &todo.CreatedAt, &todo.UpdatedAt)
	return &todo, err
}

func (r *Repository) FindByIDForUpdate(ctx context.Context, id uuid.UUID) (*domain.ToDo, error) {
	var todo domain.ToDo
	err := r.connection.QueryRow(
		ctx,
		"select id, name, body, deadline, createdAt, updatedAt from todo.todo where id=$1 for update",
		id).Scan(&todo.ID, &todo.Name, &todo.Body, &todo.Deadline, &todo.CreatedAt, &todo.UpdatedAt)
	return &todo, err
}

func (r *Repository) Save(ctx context.Context, todo *domain.ToDo) error {
	if !todo.IsNew() {
		_, err := r.connection.Exec(
			ctx,
			"update todo.todo SET name=$1, priority=$2, body=$3, deadline=$4, updatedAt=$5 where id=$6",
			todo.Name, todo.Priority, todo.Body, todo.Deadline, time.Now(), todo.ID,
		)

		return err
	}

	saveTime := time.Now()
	_, err := r.connection.Exec(
		ctx,
		"insert into todo.todo(id, priority, name, body, deadline, createdAt, updatedAt) values ($1, $2, $3, $4, $5, $6, $7)",
		todo.ID, todo.Priority, todo.Name, todo.Body, todo.Deadline, saveTime, saveTime,
	)

	return err
}
