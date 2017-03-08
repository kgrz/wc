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
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	count := wc.ReadAndCount(f)
	fmt.Println(count)
}
