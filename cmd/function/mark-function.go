package function

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/vianavitor-dev/task-tracker/models"
)

func Mark(mark string, args []string) error {

	if len(args) != 1 {
		fmt.Printf("%s %v : only one argument is supported", mark, args)
		os.Exit(0)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("%s %v : %v", mark, args, err)
	}

	var file = models.Files{PathName: "task-list.json"}
	taskList := []models.Task{}

	if err := file.FileToTasks(&taskList); err != nil && len(taskList) > 0 {
		return fmt.Errorf("%s %v : %v", mark, args, err)
	}

	if len(taskList) <= 0 {
		return fmt.Errorf("%s %v : tasks list is empty, use it only if the list has tasks", mark, args)
	}

	var found = false

	for i, t := range taskList {
		if t.ID == id {

			if mark == "markDone" {
				taskList[i].Status = "done"
			} else {
				taskList[i].Status = "in progress"
			}

			taskList[i].UpdatedAt = time.Now().UTC().Format("2006-01-02")
			found = true
		}
	}

	if !found {
		return fmt.Errorf("%s %v : task not found", mark, args)
	}

	if err := file.TruncadeTask(taskList); err != nil {
		return fmt.Errorf("%s %v : %v", mark, args, err)
	}

	return nil
}
