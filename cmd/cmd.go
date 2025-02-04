// Package cmd - package
package cmd

import (
	"errors"
	"fmt"
	"github/Asac2142/cli-tracker/file"
	"github/Asac2142/cli-tracker/task"
	"os"
	"strconv"
)

const (
	add            string = "add"
	update         string = "update"
	deleteTask     string = "delete"
	markInProgress string = "mark-in-progress"
	markDone       string = "mark-done"
	list           string = "list"
	exit           string = "exit"
)

// HandleTrackerCLI - Handles user's input
func HandleTrackerCLI(f *file.File[task.TContent]) error {
	task := task.NewTask(f)

	for {
		args := os.Args[1:]

		if len(args) == 0 {
			return errors.New("Invalid input")
		}

		switch args[0] {
		case add:
			addTask(&args, task)
		case update:
			updateTask(&args, task)
		case deleteTask:
			removeTask(&args, task)
		case markInProgress:
			updateStatus(&args, markInProgress, task)
		case markDone:
			updateStatus(&args, markDone, task)
		case list:
			listTasks(&args, task)
		case exit:
			return nil
		default:
			return errors.New("Invalid task command")
		}
	}
}

func addTask(args *[]string, t *task.Task) error {
	if len(*args) == 1 {
		return errors.New("add task requires a description")
	}

	desc := (*args)[1]
	added, err := t.Add(desc)

	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)", added.ID)

	return nil
}

func updateTask(args *[]string, t *task.Task) error {
	if len(*args) != 3 {
		return errors.New("update task requires an id & description")
	}

	strID := (*args)[1]
	id, err := strconv.Atoi(strID)
	if err != nil {
		return errors.New("invalid id provided")
	}

	dsc := (*args)[2]
	return t.Update(id, dsc)
}

func removeTask(args *[]string, t *task.Task) error {
	if len(*args) != 1 {
		return errors.New("delete task requires an id")
	}

	asString := (*args)[1]
	id, err := strconv.Atoi(asString)
	if err != nil {
		return errors.New("invalid id provided")
	}

	return t.Delete(id)
}

func updateStatus(args *[]string, status string, t *task.Task) error {
	if len(*args) != 2 {
		return errors.New("mark-in-progress task requires an id")
	}

	asString := (*args)[1]
	id, err := strconv.Atoi(asString)
	if err != nil {
		return errors.New("invalid task id provided")
	}

	var s task.Status

	if status == markInProgress {
		s = task.Inprogress
	}

	if status == markDone {
		s = task.Done
	}

	return t.UpdateStatus(id, s)
}

func listTasks(args *[]string, t *task.Task) error {
	if len(*args) == 1 {
		handlePrintTasks(nil, t)
	}

	if len(*args) != 2 {
		return errors.New("list command requires a task status")
	}

	status := task.Status((*args)[1])

	if status == task.Done {
		return handlePrintTasks(&status, t)
	}

	if status == task.Inprogress {
		return handlePrintTasks(&status, t)
	}

	if status == task.Todo {
		return handlePrintTasks(&status, t)
	}

	return errors.New("invalid task status")
}

func handlePrintTasks(s *task.Status, t *task.Task) error {
	tasks, err := t.GetByStatus(s)
	if err != nil {
		return err
	}

	for _, v := range *tasks {
		fmt.Printf("ID: %d\tDESCRIPTION: %s\tSTATUS: %s\tCREATED AT: %v\tUPDATED AT: %v\n", v.ID, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
	}

	return nil
}
