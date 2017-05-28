package main

import (
	"bufio"
	"fmt"
	"io"
	"unicode/utf8"
)

// Counts contains the snapshot of the Word, Line, Char counts after the file is
// processed.
type Counts struct {
	Words int
	Lines int
	Chars int
}

func (c Counts) String() string {
	return fmt.Sprintf("line count: %d\nword count: %d\nchar count: %d", c.Lines, c.Words, c.Chars)
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
		lineLength := len(slice)

		// Special case for empty line. Skip to the next iteration
		if lineLength == 0 {
			continue
		}

		count.Chars += utf8.RuneCount(slice)
		var isPrevCharSpace bool

		// Special case for the first character. If it's a space, then set the
		// previous char pointer to true.
		if isSpace(slice[0]) {
			isPrevCharSpace = true
		}

		// For each line, start from the second byte from the slice
		for index := 1; index < lineLength; index++ {
			char := slice[index]
			if isSpace(char) {
				if !isPrevCharSpace {
					count.Words++
				}
				isPrevCharSpace = true
			} else {
				isPrevCharSpace = false
			}
		}

		// all the bytes until the last one on a line have been counted. If the
		// previous character (last of the line) is not a space, increment the word
		// count, but only if the line has some characters.
		if !isPrevCharSpace {
			count.Words++
		}
	}

	return count
}
