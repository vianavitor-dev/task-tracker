package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/vianavitor-dev/task-tracker/models"
)

func DeleteTaskCommand() *Command {

	cmd := &Command{
		flags:   flag.NewFlagSet("delete", flag.ExitOnError),
		Execute: deleteTask,
	}

	cmd.flags.Usage = func() {
		fmt.Print(`Usage: ... delete <id>`)
	}

	return cmd
}

var deleteTask = func(c *Command, args []string) {

	if len(args) != 1 {
		fmt.Printf("deleteTask %v : only one argument is supported", args)
		os.Exit(0)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("deleteTask %v : %v", args, err)
	}

	var file = models.Files{PathName: "task-list.json"}
	taskList := []models.Task{}

	// Getting all tasks from the list
	if err := file.FileToTasks(&taskList); err != nil {
		fmt.Print(fmt.Errorf("deleteTask %v : %v", args, err))
		os.Exit(0)
	}

	if len(taskList) <= 0 {
		fmt.Printf("deleteTask %v : tasks list is empty, use it only if the list has tasks", args)
		os.Exit(0)
	}

	var found = false

	var newTaskList = make([]models.Task, len(taskList)-1)
	var index = 0

	for _, task := range taskList {
		if task.ID != id {
			if index < len(newTaskList) {
				newTaskList[index] = task
				index++
			}
		} else {
			found = true
		}
	}

	if !found {
		fmt.Printf("deleteTask %v : task not found", args)
		os.Exit(0)
	}

	if err := file.TruncadeTask(newTaskList); err != nil {
		fmt.Print(fmt.Errorf("deleteTask %v : %v", args, err))
		os.Exit(0)
	}

	fmt.Printf("Task %d successfully deleted", id)
	os.Exit(0)
}
