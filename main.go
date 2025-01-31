package main

import (
	"fmt"
	"os"

	"github.com/vianavitor-dev/task-tracker/cmd"
)

func main() {

	var c *cmd.Command

	switch os.Args[1] {
	case "add":
		c = cmd.AddTaskCommand()
	default:
		fmt.Print(fmt.Errorf("task-tracker %q is not a command", os.Args[1]))
		os.Exit(0)
	}

	c.Init(os.Args[2:])
	c.Run()
}
