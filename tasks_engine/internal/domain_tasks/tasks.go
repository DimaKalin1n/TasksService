package domain_tasks

import (
	"errors"
	"time"
)

type QueueIdType int32

type TaskStatus string

const (
	New        TaskStatus = "new"
	InProgress TaskStatus = "in_progress"
	Postponed  TaskStatus = "postponed"
	Completed  TaskStatus = "completed"
	Cancelled  TaskStatus = "cancelled"
	Queued     TaskStatus = "queued"
)

type Task struct {
	// id - Данное поле генерируется в БД
	ID            int32
	Status        TaskStatus
	Owner         string
	CompletedUser string
	CreatedAt     time.Time
	CompletedAt   *time.Time
	PostponedDate *time.Time
	Priopity      int32
	Title         string
	Description   string
	RequesterUser string
	CreateUser    string
	QueueId       QueueIdType
}

func (t *Task) GetId() int32 {
	return t.ID
}

func (t *Task) SetId(newId int32) error {
	if t.GetId() == 0 {
		t.ID = newId
		return nil
	}
	return errors.New("Bad req id")
}

func NewTask(title, desc, requestUser, createUser string, priopity int32, queueId QueueIdType) (*Task, error) {
	if title == "" || desc == "" || requestUser == "" || createUser == "" {
		return nil, ErrorBadParams
	}
	if priopity < 0 {
		return nil, ErrorPriority
	}
	if queueId == 0 {
		return nil, ErrorNotFoundQueue
	}

	return &Task{
		Status:        New,
		CreatedAt:     time.Now(),
		CreateUser:    createUser,
		Title:         title,
		Description:   desc,
		RequesterUser: requestUser,
		Priopity:      priopity,
		QueueId:       queueId,
	}, nil
}

func (t *Task) ClearOwner() error {
	if err := t.ValidTasksToUpdate(); err != nil {
		return err
	}
	t.Owner = ""
	return nil
}

func (t *Task) TaskToQueue(newQueueId QueueIdType) error {
	if err := t.ValidTasksToUpdate(); err != nil {
		return err
	}
	if newQueueId == 0 {
		return errors.New("Queue not found")
	}
	t.QueueId = newQueueId
	t.Status = Queued
	return nil
}

func (t *Task) PostponedTasks(datePostponed time.Time) error {
	if err := t.ValidTasksToUpdate(); err != nil {
		return err
	}
	t.Status = Postponed
	t.PostponedDate = &datePostponed
	return nil
}

func (t *Task) CancelledTask(user string) error {
	if err := t.ValidTasksToUpdate(); err != nil {
		return err
	}
	dateNow := time.Now()
	t.CompletedAt = &dateNow
	t.CompletedUser = user
	t.Owner = ""
	t.Status = Cancelled
	return nil
}

func (t *Task) CompletedTask(user string) error {
	if err := t.ValidTasksToUpdate(); err != nil {
		return err
	}
	timeNow := time.Now()
	t.CompletedAt = &timeNow
	t.CompletedUser = user
	t.Owner = ""
	t.Status = Completed
	return nil
}

func (t *Task) TakeInProgress(user string) error {
	if err := t.ValidTasksToUpdate(); err != nil {
		return err
	}
	if t.Status == InProgress {
		return ErrorTaskInProgres
	}
	t.Status = InProgress
	t.Owner = user
	return nil
}

func (t *Task) UpdateOwner(user string) error {
	if err := t.ValidTasksToUpdate(); err != nil {
		return err
	}
	t.Owner = user
	return nil
}

func (t *Task) ValidTasksToUpdate() error {
	if t.Status == Completed || t.Status == Cancelled {
		return ErrorTaskFinallStatus
	}
	return nil
}
