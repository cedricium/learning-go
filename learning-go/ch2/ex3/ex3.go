package main

import (
	"fmt"
	"math"
)

func main() {
	// hard-coded maximum values
	// var b byte = 255
	// var smallI int32 = 2147483647
	// var bigI uint64 = 18446744073709551615

	// using maximum constants defined in "math" package
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
