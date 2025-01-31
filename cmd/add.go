package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

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

	listFile, err := os.ReadFile("task-list.json")
	if err != nil {
		log.Fatalf("addTask %q : %v", desc, err)
	}

	taskList := []Task{}

	if err := json.Unmarshal(listFile, &taskList); err != nil && len(listFile) > 0 {
		log.Fatalf("addTask %q : %v", desc, err)
	}

	current := Task{
		ID:          len(taskList) + 1,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now().UTC().Format("2006-01-02"),
		UpdatedAt:   time.Now().UTC().Format("2006-01-02"),
	}

	taskList = append(taskList, current)
	jsonCurr, err := json.MarshalIndent(taskList, "", " ")

	if err != nil {
		fmt.Print(fmt.Errorf("addTask %q : %v", desc, err))
		os.Exit(0)
	}

	if err := os.WriteFile("task-list.json", jsonCurr, 0644); err != nil {
		log.Fatalf("addTask %q : %v", desc, err)
	}

	fmt.Printf("Task added successfully (ID: %d)", current.ID)
}
