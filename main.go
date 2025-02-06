// Package main - main package
package main

import (
	"github/Asac2142/cli-tracker/cmd"
	"github/Asac2142/cli-tracker/file"
	"github/Asac2142/cli-tracker/task"
)

// run project locally as "go run main.go"

func main() {
	f := file.New[task.TContent]()
	cmd.HandleTrackerCLI(f)
}
