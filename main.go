package main

import (
	"os"
	"fmt"
	"local/katana/cmd"
	"local/katana/internal/clitools"
	"log"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("main.go: no subcommand found")
	}
	switch os.Args[1] {
	case "split":
		cli := clitools.SplitCLIArgs{}
		cli.Init()
		if err := cmd.SplitFile(*cli.Filepath, (*cli.ChunkSize)*1024); err != nil {
			log.Fatalf("main.go: `split`:", err)
		}
	case "merge":
		cli := clitools.MergeCLIArgs{}
		cli.Init()
		if err := cmd.MergeParts(*cli.Filepath); err != nil {
			log.Fatalf("main.go: `merge`:", err)
		}
	case "check":
		if len(os.Args) != 4 {
			log.Fatalf("main.go: Invalid syntax for `check`: must be katana check <FILE_1> <FILE_2>")
		}
		stat1, _ := os.Stat(os.Args[2])
		stat2, _ := os.Stat(os.Args[3])
		fmt.Println(stat1.Size() == stat2.Size())
	default:
		log.Fatalf("main.go: unknown subcommand: %s", os.Args[1])
	}
}