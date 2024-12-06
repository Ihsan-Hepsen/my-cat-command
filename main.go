package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		if exitErr, ok := e.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}

func main() {
	numberedLines := flag.Bool("n", false, "Numbered output lines")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("Usage: my-cat [-n] <file1> [file2 ...]")
		return
	}

	for _, arg := range files {
		filePath := filepath.Join(arg)

		file, err := os.Open(filePath)
		check(err)
		defer file.Close()

		if *numberedLines {
			printWithNumbers(file)
		} else {
			_, err = io.Copy(os.Stdout, file)
			check(err)
		}

		fmt.Println()
	}

}

func printWithNumbers(file *os.File) {
	scanner := bufio.NewScanner(file)
	lineNum := 1

	for scanner.Scan() {
		fmt.Printf("%3d %s\n", lineNum, scanner.Text())
		lineNum++
	}
	check(scanner.Err())
}
