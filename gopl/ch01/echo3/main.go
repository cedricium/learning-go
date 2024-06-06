// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Exercise 1.1
	fmt.Println(os.Args[0], strings.Join(os.Args[1:], " "))

	// Exercise 1.2
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
