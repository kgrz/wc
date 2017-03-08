package wc

import (
	"bufio"
	"fmt"
	"io"
)

type Count struct {
	Words int
	Lines int
	Chars int
}

func (c Count) String() string {
	return fmt.Sprintf("word count: %d\nline count: %d\nchar count: %d", c.Words, c.Lines, c.Chars)
}

func isSpace(char byte) bool {
	return char == 32 || char == 9
}

func ReadAndCount(f io.Reader) Count {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var count Count

	for scanner.Scan() {
		count.Lines++
		count.Chars++
		slice := scanner.Bytes()
		for index, char := range slice {
			count.Chars++
			if isSpace(char) {
				if index == 0 {
					continue
				}

				previousChar := slice[index-1]
				if !isSpace(previousChar) {
					count.Words++
				}
			}
		}
		// last item is space, then don't add value
		length := len(slice)
		if length > 0 {
			lastChar := slice[length-1]
			if !isSpace(lastChar) {
				count.Words++
			}
		}
	}

	return count
}
