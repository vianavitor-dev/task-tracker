package cmd

import "flag"

type Command struct {
	flags   *flag.FlagSet
	Execute func(c *Command, args []string)
}

func (c *Command) Init(args []string) error {
	return c.flags.Parse(args)
}

func (c *Command) Called(args []string) bool {
	return c.flags.Parsed()
}

func (c *Command) Run() {
	c.Execute(c, c.flags.Args())
}
