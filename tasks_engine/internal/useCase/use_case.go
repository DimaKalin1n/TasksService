package usecase

import (
	"context"
	"time"

	domain_tasks "github.com/DimaKalin1n/TasksService/internal/domain_tasks"
)

type (
	TasksSearchStruct struct {
		Status          domain_tasks.TaskStatus
		SearchStartDate time.Time
		SearchEndDate   time.Time
		QueueId         domain_tasks.QueueIdType
		RequesterUser   string
		CompletedUser   string
	}

	TaskRepo interface {
		CreateTasks(ctx context.Context, title, desc, requestUser, createUser string, priority int32, queueId domain_tasks.QueueIdType) (*domain_tasks.Task, error)
		GetTasksById(ctx context.Context, id int32) (*domain_tasks.Task, error)
		UpdateTasks(ctx context.Context, t *domain_tasks.Task) error
		SearchTasks(ctx context.Context, t TasksSearchStruct) ([]domain_tasks.Task, error)
	}

	TaskUseCase struct {
		repo TaskRepo
	}
)

func NewTaskUseCase(repo TaskRepo) *TaskUseCase {
	return &TaskUseCase{
		repo: repo,
	}
}
