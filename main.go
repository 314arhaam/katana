package main

import (
	"os"
	"fmt"
	"strconv"
	"io"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("main.go: no subcommand found")
		return
	}
	switch os.Args[1] {
	case "split":
		//
	default:
		fmt.Println("main.go: unknown subcommand")
	}
}