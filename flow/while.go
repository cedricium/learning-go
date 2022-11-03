package main

import "fmt"

func main() {
	count := 1

	/*
		Go doesn't have a built-in `while` keyword. Instead, we can use
		a `for` with just a condition with, e.g. an incrementor to mimic
		the while functionality.
	*/
	for count <= 5 {
		fmt.Println(count)
		count += 1
	}
}
