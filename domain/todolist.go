package domain

import (
	"context"
	"time"
)

type TodoList struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type TodoListUsecase interface {
	Get(ctx context.Context) (res []TodoList, err error)
	GetByID(ctx context.Context, id int) (TodoList, error)
	Insert(ctx context.Context, t *TodoList) error
	Update(ctx context.Context, t *TodoList) error
	Delete(ctx context.Context, id int) error
}
type TodolistRepository interface {
	Get(ctx context.Context) (res []TodoList, err error)
	GetByID(ctx context.Context, id int) (TodoList, error)
	Insert(ctx context.Context, t *TodoList) error
	Update(ctx context.Context, t *TodoList) error
	Delete(ctx context.Context, id int) error
}
