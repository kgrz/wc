package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/kgrz/wc/wc"
)

func main() {
	filename := os.Args[1]
	checkIfAscii(filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	count := wc.ReadAndCount(f)
	fmt.Println(count)
}

func checkIfAscii(filename string) {
	out, err := exec.Command("file", "-0", "-b", "--mime-encoding", filename).Output()
	if err != nil {
		log.Fatal(err)
	}
	encoding := string(out)

	if encoding != "us-ascii\n" {
		log.Fatal("File encoding not supported yet. Need US-ASCII, got: ", encoding)
	}
}
