package main

import "fmt"

/* declaring single variables */
var a int

/* declaring multiple variables */
var (
	b bool
	c float32
	d string
)

func main() {
	a = 42            // single variable assignment
	b, c = true, 41.9 // multi-variable assignment
	d = "string"      // strings must use double-quotes (sorry JS-trained fingers)

	fmt.Println(a, b, c, d)
}
