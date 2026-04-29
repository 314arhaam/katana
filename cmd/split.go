package cmd

import (
	"os"
	"fmt"
	"strconv"
	"io"
	"path"
	"math"
)

func SplitFile(filepath string, chunkSize int) error {
	// get filename from path
	filename := path.Base(filepath)
	// check chunk size non zero
	if chunkSize == 0 {
		return fmt.Errorf("split.go: chunkSize cannot be zero")
	}
	// get file stat
	fileStat, err := os.Stat(filepath)
	if err != nil {
		return fmt.Errorf("split.go: stat input file:", err)	
	}
	// check if chunk size bigger than file size
	if fileStat.Size() <= int64(chunkSize) {
		return fmt.Errorf("split.go: chunkSize greater than original file size")
	}
	// check if chunk size leads to more than 50 parts
	if numParts := math.Floor(float64(fileStat.Size()) / float64(chunkSize)) + 1; numParts > 50 {
		return fmt.Errorf("split.go: chunkSize leads more than 50 parts (%.0f parts), use a greater chunkSize", numParts)
	}
	// open a file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("split.go: open input:", err)
	}
	defer file.Close()
	// chunk it
	part := int(0)
	// chunkSize := int(512*1024)
	eofOccured := false
	for !eofOccured {
		chunk := make([]byte, chunkSize)
		if num, err := file.ReadAt(chunk, int64(part*chunkSize)); err != nil{
			if err == io.EOF {
				chunk = chunk[:num]
				eofOccured = true
			} else {
				return fmt.Errorf("split.go: read input:", err)	
			}
		}
		partname := filename + "." + strconv.Itoa(int(part))
		output, err := os.Create(partname)
		if err != nil {
			return fmt.Errorf("split.go: create output:", err)
		}
		defer output.Close()
		if _, err := output.Write(chunk); err != nil {
			if err == io.EOF {
				eofOccured = true
			} else {
				return fmt.Errorf("split.go: write to output:", err)	
			}
		}
		part += 1
	}
	return nil
}