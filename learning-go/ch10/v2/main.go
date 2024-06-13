// Ch10 contains solutions to "Learning Go" book, Chapter 10.
// v2 updated to make use of Go generics.
package main

import "golang.org/x/exp/constraints"

// Number is a generic interface to allow for adding both
// ints (un- and signed) as well as floats. Backed by the
// [golang.org/x/exp/constraints.Integer] and
// [golang.org/x/exp/constraints.Float] interfaces.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two generic Numbers and returns the sum.
// Follows the rules of addition as noted here:
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
