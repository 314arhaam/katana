package clitools

import (
	"flag"
	"os"
)

type SplitCLIArgs struct {
	Filepath	*string
	ChunkSize 	*int
}

func (c *SplitCLIArgs) Init() {
	splitCmd := flag.NewFlagSet("split", flag.ExitOnError)
	c.Filepath = splitCmd.String("f", "", "Name of the file to split")
	c.ChunkSize = splitCmd.Int("s", 512, "ChunkSize in KiB")
	splitCmd.Parse(os.Args[2:])
}

type MergeCLIArgs struct {
	Filepath	*string
}

func (c *MergeCLIArgs) Init() {
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)
	c.Filepath = mergeCmd.String("f", "", "Name of the root file to merge")
	mergeCmd.Parse(os.Args[2:])
}