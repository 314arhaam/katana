package iotools

import (
	"os"
	"fmt"
)

func GetParts(filename, filedir string) ([]string, error) {
	partFiles := make([]string, 0)
	stuffsInDir, err := os.ReadDir(filedir)
	if err != nil {
		return partFiles, fmt.Errorf("fileutils.go: GetParts(...): ", error)
	}
	for _, stuff := range stuffsInDir {
		if ! stuff.IsDir() && strings.Contains(stuff.Name(), filename + ".") {
			partFiles = append(partFiles, p.Name())
		}
	}
	return partFiles, nil
}

/*
func GetFileSorter(filepath string) func(string, string) int {
	return func(a, b string){
		x, _ := strconv.Atoi(strings.Split(a, filepath + ".")[1])
		y, _ := strconv.Atoi(strings.Split(b, filepath + ".")[1])
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
		return 0
	}
}*/