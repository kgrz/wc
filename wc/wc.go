package wc

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io"
	"unicode/utf8"
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
	freqs := make(map[string]int, 0)
	h := fnv.New64()

	var count Count

	for scanner.Scan() {
		count.Lines++
		count.Chars++
		slice := scanner.Bytes()
		count.Chars += utf8.RuneCount(slice)
		lineLength := len(slice)
		var word []byte

		for index, char := range slice {
			if index == 0 {
				continue
			}

			if isSpace(char) {
				word = append(word, char)
				previousChar := slice[index-1]
				if !isSpace(previousChar) {
					h.Write(word)
					hash := fmt.Sprintf("%d", h.Sum(nil))
					freqs[hash]++
					count.Words++
				}
			} else {
				if index == lineLength-1 {
					h.Write(word)
					hash := fmt.Sprintf("%d", h.Sum(nil))
					freqs[hash]++
					count.Words++
				}

			}
		}

		h.Reset()
	}

	return count
}
