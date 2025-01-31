// Package task - Task structure
package task

import "time"

// Status - handles task status
type Status string

const (
	TODO       Status = "TODO"        // TODO - task status
	INPROGRESS Status = "IN-PROGRESS" // INPROGRESS - task status
	DONE       Status = "DONE"        // DONE - task status
)

// Task - task structure
type Task struct {
	ID          int64
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// New - creates a new task
func New(description string, status Status) *Task {
	return &Task{
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
