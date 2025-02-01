package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/vianavitor-dev/task-tracker/cmd/function"
)

func MarkDoneTaskCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("mark", flag.ExitOnError),
		Execute: markDone,
	}

	cmd.flags.Usage = func() {
		fmt.Print(`Usage: ... mark-done <id>`)
	}
	return cmd
}

var markDone = func(c *Command, args []string) {

	if err := function.Mark("markDone", args); err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	fmt.Printf("Task %q was marked successfully", args[0])
	os.Exit(0)
}
