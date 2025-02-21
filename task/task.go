// Package task - Task structure
package task

import (
	"errors"
	"github/Asac2142/cli-tracker/file"
	"slices"
	"strings"
	"time"
)

// Status - task status ("todo", "in-progress", "done")
type Status string

const (
	Todo       Status = "todo"        // Todo status
	Inprogress Status = "in-progress" // Inprogress status
	Done       Status = "done"        // Done status
)

// Task - task structure
type Task struct {
	file *file.File[TContent]
}

// TContent - task content structure
type TContent struct {
	ID          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewTask - creates a new task
func NewTask(f *file.File[TContent]) *Task {
	return &Task{file: f}
}

func newContent(description string) *TContent {
	return &TContent{
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// Add - adds a new task to json file.
func (t *Task) Add(description string) (*TContent, error) {
	var maxID int
	fileContent, err := t.file.Read()
	if err != nil {
		return nil, err
	}

	if len(fileContent) != 0 {
		maxID = slices.MaxFunc(fileContent, func(tc1, tc2 TContent) int {
			if tc1.ID > tc2.ID {
				return tc1.ID
			}
			return tc2.ID
		}).ID
	}

	id := maxID + 1
	tc := newContent(strings.TrimSpace(description))
	tc.ID = id
	fileContent = append(fileContent, *tc)

	return tc, t.file.Write(&fileContent)
}

// Update - update a task.
func (t *Task) Update(id int, dsc string) error {
	tasks, err := t.file.Read()
	if err != nil {
		return err
	}

	i := slices.IndexFunc(tasks, func(t TContent) bool { return t.ID == id })
	if i == -1 {
		return errors.New("task with id provided not found")
	}

	tc := &tasks[i]
	tc.Description = strings.TrimSpace(dsc)
	tc.UpdatedAt = time.Now()
	tasks[i] = *tc

	return t.file.Write(&tasks)
}

// Delete - deletes a task based on id.
func (t *Task) Delete(id int) error {
	tasks, err := t.file.Read()
	if err != nil {
		return err
	}

	i := slices.IndexFunc(tasks, func(task TContent) bool { return task.ID == id })
	if i == -1 {
		return errors.New("task with id provided not found")
	}

	filtered := slices.DeleteFunc(tasks, func(task TContent) bool { return task.ID == id })

	return t.file.Write(&filtered)
}

// GetByStatus - returns a list of tasks by status. If no status is provided, it returns all existing tasks.
func (t *Task) GetByStatus(s *Status) (*[]TContent, error) {
	var _tasks []TContent

	if s == nil {
		tasks, err := t.file.Read()
		if err != nil {
			return nil, err
		}

		return &tasks, nil
	}

	tasks, err := t.file.Read()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].Status == *s {
			_tasks = append(_tasks, tasks[i])
		}
	}

	return &_tasks, nil
}

// UpdateStatus - updates a task status given an id.
func (t *Task) UpdateStatus(id int, s Status) error {
	tasks, err := t.file.Read()
	if err != nil {
		return err
	}

	i := slices.IndexFunc(tasks, func(task TContent) bool { return task.ID == id })
	if i == -1 {
		return errors.New("task with id provided not found")
	}

	task := &tasks[i]
	task.Status = s
	task.UpdatedAt = time.Now()
	tasks[i] = *task

	return t.file.Write(&tasks)
}
