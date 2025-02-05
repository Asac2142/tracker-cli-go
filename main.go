// Package main - main package
package main

import (
	"github/Asac2142/cli-tracker/cmd"
	"github/Asac2142/cli-tracker/file"
	"github/Asac2142/cli-tracker/task"
)

func main() {
	f := file.New[task.TContent]()
	cmd.HandleTrackerCLI(f)
}
