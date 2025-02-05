<p align="center"><strong>Task Tracker CLI</strong></p>

## Description
Task tracker is a project used to track and manage your tasks. In this task, you will build a simple command line interface (CLI) to track what you need to do, what you have done, and what you are currently working on. This project will help you practice your programming skills, including working with the filesystem, handling user inputs, and building a simple CLI application.

## Usage 🛠
- `git clone https://github.com/Asac2142/tracker-cli-go.git`
- `cd tracker-cli-go`
- `go build -o task-cli`

## Add task ➕
```
task-cli add <task-description>
```

## Update task 🔄
```
task-cli update <task-id> <task-new-description>
```

## Delete task 🗑️
```
task-cli delete <task-id>
```

## Update task status to in-progress 🔧
```
task-cli mark-in-progress <task-id>
```
## Update task status to done ✅
```
task-cli mark-done <task-id>`
```

## Listing tasks 👀
```
task-cli list             //list all tasks
task-cli list done        //list tasks with status done
task-cli list todo        //list tasks with status todo
task-cli list in-progress //list tasks with status in-progress
```
