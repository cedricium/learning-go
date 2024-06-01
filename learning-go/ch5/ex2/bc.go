/*
Bc counts the number of bytes in the given file. Acronym for "byte count".
Essentially a simply version of `wc -c FILE`.

Usage:

	bc FILE
*/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	sum, data := 0, make([]byte, 2048)
	for {
		n, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
		sum += n
	}
	return sum, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no filename specified")
	}

	filename := os.Args[1]
	size, err := fileLen(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(size, filename)
}
