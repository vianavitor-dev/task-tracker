package function

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func Mark(mark string, args []string) {

	if len(args) != 1 {
		fmt.Printf("%s %v : only one argument is supported", mark, args)
		os.Exit(0)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("%s %v : %v", mark, args, err)
	}

	listFile, err := os.ReadFile("task-list.json")
	if err != nil {
		log.Fatalf("%s %v : %v", mark, args, err)
	}

	taskList := []Task{}

	if err := json.Unmarshal(listFile, &taskList); err != nil && len(listFile) > 0 {
		log.Fatalf("%s %v : %v", mark, args, err)
	}
	if len(taskList) <= 0 {
		fmt.Printf("%s %v : tasks list is empty, use it only if the list has tasks", mark, args)
		os.Exit(0)
	}

	var found = false

	for i, t := range taskList {
		if t.ID == id {

			if mark == "markDone" {
				taskList[i].Status = "done"
			} else {
				taskList[i].Status = "in progress"
			}

			found = true
		}
	}

	if !found {
		fmt.Printf("%s %v : task not found", mark, args)
		os.Exit(0)
	}

	jsonNewTasks, err := json.MarshalIndent(taskList, "", " ")
	if err != nil {
		fmt.Print(fmt.Errorf("%s %v : %v", mark, args, err))
		os.Exit(0)
	}

	if err := os.WriteFile("task-list.json", jsonNewTasks, 0644); err != nil {
		log.Fatalf("%s %q : %v", mark, args, err)
	}

	fmt.Printf("Task %d was marked successfully", id)
	os.Exit(0)
}
