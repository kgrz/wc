package main

import (
	"bufio"
	"fmt"
	"io"
	"unicode/utf8"
)

// Count contains the snapshot of the Word, Line, Char counts after the file is
// processed.
type Counts struct {
	Words int
	Lines int
	Chars int
}

func (c Counts) String() string {
	return fmt.Sprintf("word count: %d\nline count: %d\nchar count: %d", c.Words, c.Lines, c.Chars)
}

func isSpace(char byte) bool {
	return char == 32 || char == 9
}

// Implements the main character, word, line counting routines.
func count(f io.Reader) Counts {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var count Counts

	for scanner.Scan() {
		count.Lines++
		count.Chars++
		slice := scanner.Bytes()
		count.Chars += utf8.RuneCount(slice)
		lineLength := len(slice)

		for i := 1; i < len(slice); i++ {
			char := slice[i]
			if isSpace(char) {
				previousChar := slice[i-1]
				if !isSpace(previousChar) {
					count.Words++
				}
			} else {
				// Current character is not a space, and is the last character of the file.
				// Increment the word count.
				if i == lineLength-1 {
					count.Words++
				}
			}
		}
	}

	return count
}
