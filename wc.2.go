package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	lineCount := 0
	wordCount := 0
	charCount := 0

	for scanner.Scan() {
		lineCount++
		charCount++
		slice := scanner.Bytes()
		for _, value := range slice {
			charCount++
			if value == 32 {
				wordCount++
			}
		}
		// last item is space, then don't add value
		length := len(slice)
		if length != 0 && slice[length-1] != 32 {
			wordCount++
		}
	}

	fmt.Println("line count: ", lineCount)
	fmt.Println("char count: ", charCount)
	fmt.Println("word count: ", wordCount)
}
