// Package main - main package
package main

import (
	"fmt"
	"github/Asac2142/cli-tracker/file"
	"github/Asac2142/cli-tracker/task"
	"os"
)

func main() {
	arguments := os.Args[1:]

	fmt.Println("Arguments", arguments)

	t1 := task.New("my description 1", task.DONE)
	t2 := task.New("my description 2", task.INPROGRESS)

	t := []task.Task{*t1, *t2}
	err := file.Write(t)

	if err != nil {
		panic(err)
	}

	t, err = file.Read()
	if err != nil {
		panic(err)
	}

	for _, v := range t {
		fmt.Println(v)
	}
}
