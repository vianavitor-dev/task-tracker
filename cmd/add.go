package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/vianavitor-dev/task-tracker/models"
)

func AddTaskCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("add", flag.ExitOnError),
		Execute: addTask,
	}

	cmd.flags.Usage = func() {
		fmt.Print(`Usage: ... add "task-name"`)
	}

	return cmd
}

var addTask = func(c *Command, args []string) {
	desc := args[0]

	current := models.Task{
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now().UTC().Format("2006-01-02"),
		UpdatedAt:   time.Now().UTC().Format("2006-01-02"),
	}

	var file = models.Files{PathName: "task-list.json"}
	currentId, err := file.AppendFile(current)

	if err != nil {
		fmt.Print(fmt.Errorf("addTask %q : %v", desc, err))
	}

	fmt.Printf("Task added successfully (ID: %d)", currentId)
	os.Exit(0)
}
