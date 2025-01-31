package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

	listFile, err := os.ReadFile("task-list.json")
	if err != nil {
		log.Fatalf("deleteTask %v : %v", args, err)
	}

	taskList := []Task{}

	if err := json.Unmarshal(listFile, &taskList); err != nil && len(listFile) > 0 {
		log.Fatalf("deleteTask %v : %v", args, err)
	}
	if len(taskList) <= 0 {
		fmt.Printf("deleteTask %v : tasks list is empty, use it only if the list has tasks", args)
		os.Exit(0)
	}

	var found = false
	newTaskList := []Task{}

	for _, t := range taskList {
		if t.ID != id {
			newTaskList = append(newTaskList, t)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Printf("deleteTask %v : task not found", args)
		os.Exit(0)
	}

	jsonNewTasks, err := json.MarshalIndent(newTaskList, "", " ")
	if err != nil {
		fmt.Print(fmt.Errorf("deleteTask %v : %v", args, err))
		os.Exit(0)
	}

	if len(newTaskList) <= 0 {
		writeTask(nil, args)
	} else {
		writeTask(jsonNewTasks, args)
	}

	fmt.Printf("Task %d successfully deleted", id)
	os.Exit(0)
}

func writeTask(value []byte, args []string) {
	if err := os.WriteFile("task-list.json", value, 0644); err != nil {
		log.Fatalf("deleteTask %q : %v", args, err)
	}
}
