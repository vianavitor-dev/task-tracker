package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func UpdateTaskCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("update", flag.ExitOnError),
		Execute: updateTask,
	}

	cmd.flags.Usage = func() {
		fmt.Print(`Usage: ... update <id> <new-description>; new-description: Use "-" instead " "`)
	}

	return cmd
}

var updateTask = func(c *Command, args []string) {

	if len(args) < 2 {
		fmt.Printf("updateTask %v : arguments not found, required 2 arguments", args)
		os.Exit(0)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("updateTask %v : %v", args, err)
	}

	newDesc := strings.ReplaceAll(args[1], "-", " ")

	listFile, err := os.ReadFile("task-list.json")
	if err != nil {
		log.Fatalf("updateTask %v : %v", args, err)
	}

	taskList := []Task{}

	if err := json.Unmarshal(listFile, &taskList); err != nil && len(listFile) > 0 {
		log.Fatalf("updateTask %v : %v", args, err)
	}
	if len(taskList) <= 0 {
		fmt.Printf("updateTask %v : tasks list is empty, use it only if the list has tasks", args)
		os.Exit(0)
	}

	var found = false

	for i, t := range taskList {
		if t.ID == id {

			taskList[i].Description = newDesc
			taskList[i].UpdatedAt = time.Now().UTC().Format("2006-01-02")

			found = true
			break
		}
	}

	if !found {
		fmt.Printf("updateTask %v : task not found", args)
		os.Exit(0)
	}

	jsonTaskList, err := json.MarshalIndent(taskList, "", " ")
	if err != nil {
		fmt.Print(fmt.Errorf("updateTask %v : %v", args, err))
		os.Exit(0)
	}

	if err := os.WriteFile("task-list.json", jsonTaskList, 0644); err != nil {
		log.Fatalf("updateTask %v : %v", args, err)
	}

	fmt.Printf("Task %d, %q, successfully modified", id, newDesc)
	os.Exit(0)
}
