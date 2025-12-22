package usecase

import (
	"context"
	"time"

	domain_tasks "github.com/DimaKalin1n/TasksService/internal/domain_tasks"
)

type (
	TasksSearchStruct struct {
		Status          *domain_tasks.TaskStatus
		SearchStartDate *time.Time
		SearchEndDate   *time.Time
		QueueId         *domain_tasks.QueueIdType
		RequesterUser   *string
		CompletedUser   *string
	}

	TaskRepo interface {
		SaveTask(ctx context.Context, task *domain_tasks.Task) error
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

func (usc TaskUseCase) CreateTasks(ctx context.Context, title, desc, creatUser, requestUser string, priop int32, queueId domain_tasks.QueueIdType) (*domain_tasks.Task, error) {
	task, err := domain_tasks.NewTask(
		title, desc, requestUser, creatUser, priop, queueId,
	)
	if err != nil {
		return nil, err
	}

	if err := usc.repo.SaveTask(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (usc TaskUseCase) CompletedTasks(ctx context.Context, id int32, user string) error {
	task, err := usc.repo.GetTasksById(ctx, id)
	if err != nil {
		return err
	}
	if errCompl := task.CompletedTask(user); err != nil {
		return errCompl
	}

	if errComplTasks := usc.repo.UpdateTasks(ctx, task); errComplTasks != nil {
		return errComplTasks
	}
	return nil

}

func (usc TaskUseCase) CancellTasks(ctx context.Context, id int32, user string) error {
	task, err := usc.repo.GetTasksById(ctx, id)
	if err != nil {
		return err
	}
	if errCompl := task.CancelledTask(user); err != nil {
		return errCompl
	}

	if errComplTasks := usc.repo.UpdateTasks(ctx, task); errComplTasks != nil {
		return errComplTasks
	}
	return nil

}

func (usc TaskUseCase) ClearOwnerOnTasks(ctx context.Context, id int32) error {
	task, err := usc.repo.GetTasksById(ctx, id)
	if err != nil {
		return nil
	}

	errOwner := task.ClearOwner()
	if errOwner != nil {
		return errOwner
	}
	return nil
}

func (usc TaskUseCase) TasksToQueue(ctx context.Context, id int32, queueId domain_tasks.QueueIdType) error {
	task, err := usc.repo.GetTasksById(ctx, id)
	if err != nil {
		return err
	}
	if errToQueue := task.TaskToQueue(queueId); errToQueue != nil {
		return errToQueue
	}

	return nil
}
