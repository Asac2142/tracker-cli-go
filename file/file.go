// Package file - Handle file operations
package file

import (
	"encoding/json"
	"github/Asac2142/cli-tracker/task"
	"os"
)

const filename = "tasks.json"

// Write - writes tasks into a JSON file
func Write(data []task.Task) error {
	asJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, asJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Read - reads tasks json file
func Read() ([]task.Task, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	bytes, err := os.ReadFile(jsonFile.Name())
	if err != nil {
		return nil, err
	}

	var tasks []task.Task

	err = json.Unmarshal(bytes, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
