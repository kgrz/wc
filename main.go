package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kgrz/wc/wc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Wrong number of arguments. Basic usage: go run main.go <filename>")
	}

	if len(os.Args) > 2 {
		fmt.Printf("Warning: Multile files are not supported yet. Using the first one.\n\n")
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	counts := wc.Count(f)
	fmt.Println(counts)
}
