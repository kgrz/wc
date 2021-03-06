package main

import (
	"strings"
	"testing"
)

type fixtures struct {
	input string
	words int
	lines int
	chars int
}

// Each of the test case results below are based on the output of `wc` program
// on a plain text file with the string contents. I'm trading off actual
// correctness in favour of verifiable comparision between wc and my wc program.
// For example, one would expect a hyphenated word to count as one, but the
// *nix wc program counts them as two words.
var tests = []fixtures{
	{"this is a test string", 5, 1, 22},
	{"this   is          a     test string", 5, 1, 37},
	{"this	is			a		test string", 5, 1, 25},
	{"this is a test string\nthis is another string", 9, 2, 45},
	{"this is a test string\n\n\n\n\n\nthis is another string", 9, 7, 50},
	{"this is a test string\n this is another string", 9, 2, 46},
	{"this is a test str-\ning this is another string", 10, 2, 47},
	{"this is a test string \n this is another string", 9, 2, 47},
	{"this * * ***** is a test string", 8, 1, 32},
	{". . . . . . . . . .", 10, 1, 20},
	{`Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
	tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
	vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,
	no sea takimata sanctus est Lorem ipsum dolor sit amet.`, 50, 4, 299},
	{"é", 1, 1, 2},
	{"😀", 1, 1, 2},
	{"👂🏼", 1, 1, 3},
	{"😀😀\n👂🏼", 2, 2, 6},
	{"😀 😀\n👂🏼", 3, 2, 7},
}

func TestCount(t *testing.T) {
	for i, fixture := range tests {
		count := count(strings.NewReader(fixture.input))

		if count.Words != fixture.words {
			t.Error(
				"For test number", i+1,
				"expected", fixture.words,
				"words but",
				"got", count.Words,
				"words",
			)
		}

		if count.Lines != fixture.lines {
			t.Error(
				"For test number", i+1,
				"expected", fixture.lines,
				"lines but",
				"got", count.Lines,
				"lines",
			)
		}

		if count.Chars != fixture.chars {
			t.Error(
				"For test number", i+1,
				"expected", fixture.chars,
				"characters but",
				"got", count.Chars,
				"characters",
			)
		}
	}
}

func BenchmarkCount(b *testing.B) {
	input := ` Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed
	diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
	erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et
	ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem
	ipsum dolor sit amet.

	Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
	tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
	vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,
	no sea takimata sanctus est Lorem ipsum dolor sit amet.

	diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
	erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et
	ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem
	ipsum dolor sit amet.

	Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
	tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
	vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,
	no sea takimata sanctus est Lorem ipsum dolor sit amet.

	diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
	erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et
	ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem
	ipsum dolor sit amet.

	Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
	tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
	vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,
	no sea takimata sanctus est Lorem ipsum dolor sit amet.

	diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
	erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et
	ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem
	ipsum dolor sit amet.

	Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
	tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
	vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,
	no sea takimata sanctus est Lorem ipsum dolor sit amet.

	diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
	erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et
	ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem
	ipsum dolor sit amet.

	Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
	tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
	vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,
	no sea takimata sanctus est Lorem ipsum dolor sit amet.

	diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam
	erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et
	ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem
	ipsum dolor sit amet.

	Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
	tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
	vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,
	no sea takimata sanctus est Lorem ipsum dolor sit amet.
	`
	reader := strings.NewReader(input)

	for i := 0; i < b.N; i++ {
		count(reader)
	}
}
