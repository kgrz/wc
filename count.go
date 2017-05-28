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

		for index, char := range slice {
			if index == 0 {
				continue
			}

			if isSpace(char) {
				previousChar := slice[index-1]
				if !isSpace(previousChar) {
					count.Words++
				}
			} else {
				if index == lineLength-1 {
					count.Words++
				}

			}
		}
	}

	return count
}
