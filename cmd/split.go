package cmd

import (
	"os"
	"fmt"
	"strconv"
	"io"
)

func main() {
	// open a file
	file, err := os.Open("yeah_0-524290.txt")
	if err != nil {
		fmt.Println("main.go: open input:", err)
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
				fmt.Println("main.go: read input:", err)
				return	
			}
		}
		filename := "part-" + strconv.Itoa(int(part)) + ".katana"
		output, err := os.Create(filename)
		if err != nil {
			fmt.Println("main.go: create output:", err)
			return
		}
		defer output.Close()
		if _, err := output.Write(chunk); err != nil {
			if err == io.EOF {
				eofOccured = true
			} else {
				fmt.Println("main.go: write to output:", err)
				return	
			}
		}
		part += 1
	}
}