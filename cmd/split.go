package cmd

import (
	"os"
	"fmt"
	"strconv"
	"io"
)

func SplitFile(filepath string) error {
	// open a file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("split.go: open input:", err)
	}
	defer file.Close()
	// chunk it
	part := int(0)
	chunkSize := int(512*1024)
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
		filename := "part-" + strconv.Itoa(int(part)) + ".katana"
		output, err := os.Create(filename)
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