package main

import "fmt"

func main() {
	a := 42            // single variable initialization + assignment
	b, c := true, 41.9 // multi-variable initialization + assignment
	d := "string"      // strings still use double-quotes

	fmt.Println(a, b, c, d)
}
