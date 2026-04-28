package clitools

import (
	"flag"
	"strings"
	"os"
)

type SplitCLIArgs struct {
	Filepath	*string
	Filename	string
}

func (c *SplitCLIArgs) Init() {
	splitCmd := flag.NewFlagSet("split", flag.ExitOnError)
	c.Filepath = splitCmd.String("f", "", "Name of the file to split")
	splitCmd.Parse(os.Args[2:])
	subs := strings.Split(*c.Filepath, "/")
	c.Filename = subs[len(subs)-1]
}

type MergeCLIArgs struct {}

func (c *MergeCLIArgs) Init() {}