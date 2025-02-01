package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/vianavitor-dev/task-tracker/cmd/function"
)

func MarkInProgressTaskCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("mark", flag.ExitOnError),
		Execute: markInProgress,
	}

	cmd.flags.Usage = func() {
		fmt.Print(`Usage: ... mark-in-progress <id>`)
	}
	return cmd
}

var markInProgress = func(c *Command, args []string) {

	if err := function.Mark("markInProgress", args); err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	fmt.Printf("Task %q was marked successfully", args[0])
	os.Exit(0)
}
