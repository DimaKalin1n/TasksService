package postgres

import (
	"database/sql"
	"time"
)

type (
	TaskModel struct {
		ID            int32
		Status        string
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
		QueueId       int32
	}

	PostgresTaskRepo struct {
		db *sql.DB
	}
)
