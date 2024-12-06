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

	for _, arg := range os.Args[1:] {
		filePath := filepath.Join(arg)

		file, err := os.Open(filePath)
		check(err)
		defer file.Close()

		_, err = io.Copy(os.Stdout, file)
		check(err)
		fmt.Println()
	}

}
