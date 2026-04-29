package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"
	"slices"
	"strconv"
	"io"
)

func MergeParts(filepath string){
	filedir := path.Dir(filepath)
	filename := path.Base(filepath)
	parts, _ := os.ReadDir(filedir)
	partFiles := make([]string, 0)
	for _, p := range parts {
		if ! p.IsDir() && strings.Contains(p.Name(), filename + ".") {
			partFiles = append(partFiles, p.Name())
		}
	}
	slices.SortFunc(
		partFiles,
		func(a, b string){
			x, _ := strconv.Atoi(strings.Split(a, filepath + ".")[1])
			y, _ := strconv.Atoi(strings.Split(b, filepath + ".")[1])
			if x > y {
				return 1
			} else if x < y {
				return -1
			}
			return 0
		}
	)
	file, _ := os.Create("katana_" + filename)
	defer file.Close()
	chunkSize := 512*1024
	for i, pn := range partFiles {
		pf := path.Join(filedir, pn)
		pff, _ := os.Open(pf)
		data, _ := io.ReadAll(pff)
		file.WriteAt(data, int64(i*chunkSize))
		pff.Close()
	}
}