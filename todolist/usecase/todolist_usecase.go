package usecase

import (
	"context"
	"go_clean/domain"
	"time"
)

type todoListUseCase struct {
	todoListRepo   domain.TodolistRepository
	contextTimeout time.Duration
}

func NewTodoListUseCase(tdr domain.TodolistRepository, timeout time.Duration) domain.TodoListUsecase {
	return &todoListUseCase{
		todoListRepo:   tdr,
		contextTimeout: timeout,
	}
}

func (tdu *todoListUseCase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, tdu.contextTimeout)
	defer cancel()

	_, err := tdu.todoListRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = tdu.todoListRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (tdu *todoListUseCase) Update(ctx context.Context, t *domain.TodoList) error {
	ctx, cancel := context.WithTimeout(ctx, tdu.contextTimeout)
	defer cancel()

	err := tdu.todoListRepo.Update(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func (tdu *todoListUseCase) Insert(ctx context.Context, t *domain.TodoList) error {
	ctx, cancel := context.WithTimeout(ctx, tdu.contextTimeout)
	defer cancel()

	err := tdu.todoListRepo.Insert(ctx, t)
	return err
}

func (tdu *todoListUseCase) Get(ctx context.Context) (res []domain.TodoList, err error) {
	ctx, cancel := context.WithTimeout(ctx, tdu.contextTimeout)
	defer cancel()

	res, err = tdu.todoListRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (tdu *todoListUseCase) GetByID(ctx context.Context, id int) (res domain.TodoList, err error) {
	ctx, cancel := context.WithTimeout(ctx, tdu.contextTimeout)
	defer cancel()

	resTodoList, err := tdu.todoListRepo.GetByID(ctx, id)
	if err != nil {
		return resTodoList, err
	}
	return resTodoList, nil
}
