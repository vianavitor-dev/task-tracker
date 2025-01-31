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
		fmt.Print(`Usage: ... mark-done <id>`)
	}
	return cmd
}

var markInProgress = func(c *Command, args []string) {

	function.Mark("markInProgress", args)
	os.Exit(0)
}
