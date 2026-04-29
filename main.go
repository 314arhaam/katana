package main

import (
	"os"
	"fmt"
	"local/katana/cmd"
	"local/katana/internal/clitools"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("main.go: no subcommand found")
		return
	}
	switch os.Args[1] {
	case "split":
		cli := clitools.SplitCLIArgs{}
		cli.Init()
		if err := cmd.SplitFile(*cli.Filepath); err != nil {
			fmt.Println("main.go: `split`:", err)
			return
		}
	case "merge":
		cli := clitools.MergeCLIArgs{}
		cli.Init()
		if err := cmd.MergeParts(*cli.Filepath); err != nil {
			fmt.Println("main.go: `merge`:", err)
			return
		}
	default:
		fmt.Println("main.go: unknown subcommand")
	}
}