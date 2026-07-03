package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// define chunk size in bytes= 4KB = 4*1024 = 4096
	const chunkSize = 51200
	buffer := make([]byte, chunkSize)

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		// Process the chunk (only n bytes are valid)
		fmt.Printf("Process the chunks size %d\n", n)
		// Example: print first 50 bytes of the chunk
		fmt.Println(string(buffer))
	}

}
