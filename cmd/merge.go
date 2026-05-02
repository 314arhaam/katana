package cmd

import (
	"fmt"
	"os"
	"path"
	"slices"
	"io"
	"local/katana/internal/iotools"
)

func MergeParts(filepath string) error {
	chunkSize := int64(0)
	filedir := path.Dir(filepath)
	filename := path.Base(filepath)
	filePartNames, err := iotools.GetParts(filename, filedir)
	if err != nil {
		return fmt.Errorf("merge.go: MergeParts(...): ", err)
	}
	slices.SortFunc(
		filePartNames,
		iotools.SortParts(filename),
	)
	file, err := os.Create("katana_" + filename)
	if err != nil {
		return fmt.Errorf("merge.go: MergeParts(...): create output file: ", err)
	}
	defer file.Close()
	for i, partName := range filePartNames {
		if i == 0 {
				fileinfo, err := os.Stat(path.Join(filedir, partName))
				if err != nil {
						return fmt.Errorf("merge.go: MergeParts(...): stat partition file: ", err)
				}
				chunkSize = fileinfo.Size()
		}
		partFile, err := os.Open(path.Join(filedir, partName))
		if err != nil {
			return fmt.Errorf("merge.go: MergeParts(...): open partition file: ", err)
		}
		data, err := io.ReadAll(partFile)
		partFile.Close()
		if err != nil {
			return fmt.Errorf("merge.go: MergeParts(...): read partition file: ", err)
		}
		if _, err := file.WriteAt(data, int64(i)*chunkSize); err != nil {
			return fmt.Errorf("merge.go: MergeParts(...): write partition data to output: ", err)
		}
	}
	return nil
}