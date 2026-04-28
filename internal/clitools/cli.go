package clitools

import (
	"flag"
	"string"
)

type SplitCLIArgs struct {
	Filepath	*string
	Filename	string
}

func (c *SplitCLIArgs) Init() {
	c.Filepath = flag.String("-f", "", "Name of the file to split")
	subs := strings.Split(*c.Filepath, "/")
	c.Filename = subs[len(subs)-1]
}

type MergeCLIArgs struct {}

func (c *MergeCLIArgs) Init() {}