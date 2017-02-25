package main

import (
	"fmt"
	"log"
	"os"
)

const BYTES_TO_READ = 10000

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var charCount int
var lineCount int
var wordCount int

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	assert(err)
	defer f.Close()

	buffer := make([]byte, BYTES_TO_READ)
	stat, err := f.Stat()
	assert(err)
	fileSize := int(stat.Size())
	var lastChar byte
	// Initialize this variable to something that's not expected to be in
	// normal text file. Bad assumption, but let's do this for now
	lastChar = 4

	iterationsRequired := (fileSize / BYTES_TO_READ) + 1

	for i := 0; i < iterationsRequired; i++ {
		bytesRead, err := f.Read(buffer)
		assert(err)
		// Wont work with utf-8 characters
		charCount += bytesRead
		for j := 0; j < bytesRead; j++ {
			char := buffer[j]

			if char == 10 {
				lineCount++
			}

			// handle special cases where the word might be split in seams
			// between consecutive buffers. Rule engine is as follows:
			//
			// previous_buffer_char		current_char	increment
			//
			// alphabet					alphabet		after first space
			// alphabet					new line		immediately
			// alphabet					space			immediately
			// space					alphabet		after first space
			// space					new line		skip
			// space					space			skip
			// new line					alphabet		after first space
			// new line					new line		skip
			// new line					space			skip
			if char == 32 || char == 10 {
				if j == 0 {
					if lastChar == 32 || lastChar == 10 {
						continue
					} else {
						wordCount++
						continue
					}
				}

				wordCount++
			}

			if j == bytesRead-1 {
				// last byte
				lastChar = char
			}
		}
	}

	fmt.Println("char count: ", charCount)
	fmt.Println("word count: ", wordCount)
	fmt.Println("line count: ", lineCount)
}
