package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"
	"slices"
	"strconv"
	"io"
	"local/katana/internal/iotools"
)

func MergeParts(filepath string) error {
	filedir := path.Dir(filepath)
	filename := path.Base(filepath)
	filePartNames, err := iotools.GetParts(filename, filedir)
	slices.SortFunc(
		filePartNames,
		func(a, b string) int {
			x, _ := strconv.Atoi(strings.Split(a, filepath + ".")[1])
			y, _ := strconv.Atoi(strings.Split(b, filepath + ".")[1])
			if x > y {
				return 1
			} else if x < y {
				return -1
			}
			return 0
		},
	)
	file, err := os.Create("katana_" + filename)
	if err != nil {
		return fmt.Errorf("merge.go: MergeParts(...): create output file: ", err)
	}
	defer file.Close()
	chunkSize := 512*1024
	for i, partName := range filePartNames {
		partFile, err := os.Open(path.Join(filedir, partName))
		if err != nil {
			return fmt.Errorf("merge.go: MergeParts(...): open partition file: ", err)
		}
		data, err := io.ReadAll(partFile)
		if err != nil {
			return fmt.Errorf("merge.go: MergeParts(...): read partition file: ", err)
		}
		if _, err := file.WriteAt(data, int64(i*chunkSize)); err != nil {
			fmt.Errorf("merge.go: MergeParts(...): write partition data to output: ", err)
		}
		partFile.Close()
	}
	return nil
}