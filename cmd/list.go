package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
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

	listFile, err := os.ReadFile("task-list.json")
	if err != nil {
		log.Fatalf("addTask %v : %v", args, err)
	}

	if len(args) <= 0 {
		fmt.Printf("%s", listFile)
	} else {

		status := args[0]
		taskList := []Task{}

		if err := json.Unmarshal(listFile, &taskList); err != nil && len(listFile) > 0 {
			log.Fatalf("addTask %q : %v", status, err)
		}

		result := []Task{}

		for _, t := range taskList {
			if t.Status == status {
				result = append(result, t)
			}
		}

		jsonResult, err := json.MarshalIndent(result, "", " ")
		if err != nil {
			fmt.Print(fmt.Errorf("addTask %q : %v", status, err))
			os.Exit(0)
		}

		fmt.Printf("%s\n", jsonResult)
	}

	os.Exit(0)
}
