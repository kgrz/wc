package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
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
		for index, char := range slice {
			charCount++
			// Treat tabs as spaces too.
			if isSpace(char) {
				if index == 0 {
					continue
				}

				previousChar := slice[index-1]
				if !isSpace(previousChar) {
					wordCount++
				}
			}
		}
		// last item is space, then don't add value
		length := len(slice)
		if length == 0 {
			continue
		}

		lastChar := slice[length-1]
		if !isSpace(lastChar) {
			wordCount++
		}
	}

	fmt.Println("word count: ", wordCount)
	fmt.Println("line count: ", lineCount)
	fmt.Println("char count: ", charCount)
}

func isSpace(char byte) bool {
	return char == 32 || char == 9
}
