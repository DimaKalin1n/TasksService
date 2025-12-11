package domain

import (
	"context"
	"errors"
	"time"
)

type StatusTask string

const (
	New         StatusTask = "new"
	In_Progress StatusTask = "in_Progress"
	Cancelled   StatusTask = "cancelled"
	Completed   StatusTask = "completed"
	Postponed   StatusTask = "postponed"
)

type Task struct {
	ID          string     `json:"id"`
	AuthorTasks string     `json:"author"`
	OwnerTasks  string     `json:"owner,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      StatusTask `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
	Priority    int        `json:"priority"`
}

type TaskUseCase interface {
	TasksCreated(ctx context.Context, user, title, description string, priority int) (int, error)
	TaskInfo(ctx context.Context, id string) (Task, error)
	ChangeStatusTasks(ctx context.Context, id string) error
	ValidOperation(ctx context.Context, id string) error
	PostponedTasks(ctx context.Context, id string, time time.Time) error
	AddNoteToTasks(ctx context.Context, id, text string) error
}

type TasksRepo interface {
	Create(ctx context.Context, t *Task) (int, error)
	Update(ctx context.Context, t *Task) error
	GetById(ctx context.Context, id string) (*Task, error)
	GetListId(ctx context.Context) ([]Task, error)
}

var (
	ErrInvalidStatus    = errors.New("Некорректный статус")
	ErrTasksNotFound    = errors.New("Задание не найдено")
	ErrTasksFinalStatus = errors.New("Задание в финальном статусе изменение невозможно")
	ErrBadRequests      = errors.New("Некорректный запрос")
)
