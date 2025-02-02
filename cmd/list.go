package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vianavitor-dev/task-tracker/models"
)

func ListTaksCommand() *Command {

	cmd := &Command{
		flags:   flag.NewFlagSet("list", flag.ExitOnError),
		Execute: listTasks,
	}

	cmd.flags.Usage = func() {
		fmt.Print(`Usage: ... list "status[done, todo, in-progress]" or ""`)
	}

	return cmd
}

var listTasks = func(c *Command, args []string) {

	var fs = models.Files{PathName: "task-list.json"}

	if len(args) <= 0 {
		files, err := fs.ReadFile()
		if err != nil {
			log.Fatalf("listTasks %v : %v", args, err)
		}

		fmt.Printf("%s", files)

	} else {

		status := strings.ReplaceAll(args[0], "-", " ")
		var tasks = []models.Task{}

		if err := fs.FileToTasks(&tasks); err != nil {
			log.Fatalf("listTasks %v : %v", args, err)
		}

		result := []models.Task{}

		for _, t := range tasks {
			if t.Status == status {
				result = append(result, t)
			}
		}

		jsonResult, err := json.MarshalIndent(result, "", " ")
		if err != nil {
			fmt.Print(fmt.Errorf("listTasks %q : %v", status, err))
			os.Exit(0)
		}

		fmt.Printf("%s\n", jsonResult)
	}

	os.Exit(0)
}
