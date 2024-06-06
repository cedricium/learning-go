// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
//
// Modified to complete Exercise 1.4
package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/maps"
)

type Set = map[string]bool

type LineCount struct {
	Lines map[string]int // Lines counts the number of repeated lines
	Files map[string]Set // Files is a set of filenames where each line is found
}

func main() {
	lc := &LineCount{
		Lines: make(map[string]int),
		Files: make(map[string]Set),
	}

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, lc)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, lc)
			f.Close()
		}
	}
	for line, n := range lc.Lines {
		if n > 1 {
			fmt.Printf("%d\t%v\t%v\n", n, line, maps.Keys(lc.Files[line]))
		}
	}
}

func countLines(f *os.File, lc *LineCount) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		lc.Lines[line]++
		if _, ok := lc.Files[line]; !ok {
			lc.Files[line] = make(Set, 0)
		}
		lc.Files[line][f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
