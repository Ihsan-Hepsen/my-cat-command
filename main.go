package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: my-cat <file name>")
		return
	}

	fileName := os.Args[1]
	filePath := filepath.Join(fileName)

	// read file
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	check(err)

	defer fmt.Println()
}
