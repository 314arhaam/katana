package clitools

import (
	"flag"
	"os"
)

type SplitCLIArgs struct {
	Filepath	*string
}

func (c *SplitCLIArgs) Init() {
	splitCmd := flag.NewFlagSet("split", flag.ExitOnError)
	c.Filepath = splitCmd.String("f", "", "Name of the file to split")
	splitCmd.Parse(os.Args[2:])
}

type MergeCLIArgs struct {}

func (c *MergeCLIArgs) Init() {}