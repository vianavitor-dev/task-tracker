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

	function.Mark("markDone", args)
	os.Exit(0)
}
