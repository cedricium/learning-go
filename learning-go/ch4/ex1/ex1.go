package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// s := []int{}             // valid, but not most memory-efficient
	s := make([]int, 0, 100) // better since we know capacity of slice upfront
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(101))
	}
	fmt.Println(s)
}
