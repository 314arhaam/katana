package main

import (
	"os"
	"fmt"
	"local/katana/cmd"
	"local/katana/internal/clitools"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "main.go: no subcommand found")
		return
	}
	switch os.Args[1] {
	case "split":
		cli := clitools.SplitCLIArgs{}
		cli.Init()
		if err := cmd.SplitFile(*cli.Filepath, (*cli.ChunkSize)*1024); err != nil {
			fmt.Fprintln(os.Stderr, "main.go: `split`:", err)
			return
		}
	case "merge":
		cli := clitools.MergeCLIArgs{}
		cli.Init()
		if err := cmd.MergeParts(*cli.Filepath); err != nil {
			fmt.Fprintln(os.Stderr, "main.go: `merge`:", err)
			return
		}
	case "check":
		if len(os.Args) != 4 {
			fmt.Fprintln(os.Stderr, "main.go: Invalid syntax for `check`: must be katana check <FILE_1> <FILE_2>")
			return
		}
		stat1, _ := os.Stat(os.Args[2])
		stat2, _ := os.Stat(os.Args[3])
		fmt.Println(stat1.Size() == stat2.Size())
	default:
		fmt.Fprintln(os.Stderr, "main.go: unknown subcommand")
		return
	}
}