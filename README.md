A Shoddy Word count program
============================

Run `go run main.go <filename>`


The current version is multibyte aware. Compare the output from this
program to that of `wc -wlm <filename>`. The `-wlm` arguments output the
words, lines, characters out of the file. By default, the `wc` command
prints out words, lines, and bytes. For words/characters that support
multi-byte lengths, the number of bytes will be different from the
number of characters. For instance, the letter `é` is the length of 2
bytes, but it's considered to be a single character width (codepoint, in
UTF-8 lingo).


Testing
=======

I've added some rudimentary tests for the main counting routing in
`wc/wc_test.go`. Run the tests using `go test wc/wc_test.go`.
