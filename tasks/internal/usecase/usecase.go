package usecase

import (
	"context"
	"errors"
	"tasks/internal/domain"
	"time"
)

type TasksUseCaseInterface interface {
	TasksCreated(ctx context.Context, user, title, description string, priority int) (int, error)
}

type taskUseCase struct {
	repo domain.TasksRepo
}

func (tuc *taskUseCase) TasksCreated(ctx context.Context, user, title, description string, priority int) (int, error) {
	if title == "" {
		return 0, errors.New("Заголовой не может быть пустым")
	} else if user == "" {
		return 0, errors.New("Пользователь не может быть пустым")
	}
	t := &domain.Task{
		Title:       title,
		AuthorTasks: user,
		Description: description,
		CreatedAt:   time.Now(),
		Priority:    priority,
		Status:      domain.New,
	}
	id, err := tuc.repo.Create(ctx, t)
	if err != nil {
		return 0, errors.New("Ошибка при создании задания")
	}
	return id, nil
}
