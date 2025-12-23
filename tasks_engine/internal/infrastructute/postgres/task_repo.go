package postgres

import (
	"context"
	"database/sql"

	"github.com/DimaKalin1n/TasksService/internal/domain_tasks"
)

func NewPostgresRepo(db *sql.DB) *PostgresTaskRepo {
	return &PostgresTaskRepo{
		db: db,
	}
}

func (ptr PostgresTaskRepo) SaveTask(ctx context.Context, task *domain_tasks.Task) error {
	if task.GetId() == 0 {
		ptr.insert(ctx, task)
	}
	return ptr.update(ctx, task)
}

func (ptr *PostgresTaskRepo) insert(ctx context.Context, task *domain_tasks.Task) error {
	query := `
	INSERT INTO tasks(
		title, 
		description,
		status,
		priority,
		queue_id,
		request_user,
		create_user,
		created_at
	)
		
	VALUES( $1,$2, $3, $4, $5, $6, &7, $8)
	RETURNING id
	`
	var id int32

	err := ptr.db.QueryRowContext(ctx, query, task.Title, task.Description, task.Status, task.Priopity, task.QueueId, task.RequesterUser, task.CreateUser, task.CreatedAt).Scan(&id)
	if err != nil {
		return err
	}
	task.SetId(id)
	return nil
}

func (ptr *PostgresTaskRepo) update(ctx context.Context, task *domain_tasks.Task) error {
	query := `
		UPDATE tasks
		SET
			status = $1,
			priority = $2,
			queue_id, = $3,
			completed_user = $4,
			completed_at = $5,
			updated_at = $6,
			requester_user = $7
		WHERE id = $8
	`
	_, err := ptr.db.ExecContext(ctx, query, task.Status, task.Priopity, task.QueueId, task.CompletedUser, task.CompletedAt, task.RequesterUser, task.GetId())
	if err != nil {
		return err
	}
	return nil
}
