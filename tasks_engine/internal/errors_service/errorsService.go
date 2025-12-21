package errors_service

import "errors"

var (
	ErrorBadParams        = errors.New("Bad params to created task")
	ErrorPriority         = errors.New("Priority < 0, bad requests")
	ErrorTaskFinallStatus = errors.New("Task in finel status")
	ErrorTaskInProgres    = errors.New("Task In Progres Now")
	ErrorNotFoundQueue    = errors.New("Queue not found")
)
