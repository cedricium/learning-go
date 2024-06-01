package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}

/* OUTPUT:

0
1
2
3
4
5
6
7
8
9
0
*/

/* REASON:
The bug is on line 8, with the use of `:=` rather than `=`. When initialized,
`total` is set to the zero value for an int, or 0. Inside the loop, `total` is
being shadowed, always referencing the outer, first-initialized value (0)
rather than the expected sum of `total` and the current value of `i`.

Replacing `:=` with `=` would correctly update the true value of `total` as
intended, and the summation of the first 10 digits starting at 0 would be
displayed in the final print statement:

0
1
3
6
10
15
21
28
36
45
45
*/
