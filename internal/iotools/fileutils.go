package iotools

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func GetParts(filename, filedir string) ([]string, error) {
	partFiles := make([]string, 0)
	stuffsInDir, err := os.ReadDir(filedir)
	if err != nil {
		return partFiles, fmt.Errorf("fileutils.go: GetParts(...): ", err)
	}
	for _, stuff := range stuffsInDir {
		if ! stuff.IsDir() && strings.Contains(stuff.Name(), filename + ".") {
			partFiles = append(partFiles, stuff.Name())
		}
	}
	if len(partFiles) == 0 {
		return partFiles, fmt.Errorf("fileutils.go: GetParts(...): No partitions found for file: ", filename)
	}
	return partFiles, nil
}

func SortParts(filepath string) func(string, string) int {
	return func(a, b string) int {
		x, err := strconv.Atoi(strings.Split(a, filepath + ".")[1])
		if err != nil {
			panic("Error in SortParts")
		}
		y, err := strconv.Atoi(strings.Split(b, filepath + ".")[1])
		if err != nil {
			panic("Error in SortParts")
		}
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
		return 0
	}
}