package main

import (
	"fmt"
	"reflect"
)

func main() {
	foods := [8]string{
		"mango", "kiwi", "strawberry",
		"spinach", "broccoli", "asparagus",
		"Reese's", "M&Ms",
	}

	/* defining slices */
	fruits := foods[0:3] // simple [<include>:<exclude>] syntax

	veggies := make([]string, 3) // initializing and allocating slice with size of 3
	for i := 3; i < 6; i++ {
		veggies[i-3] = foods[i]
	}

	candy := foods[6:] // [<include>:__] syntax, similar to Python

	fmt.Printf("fruits: %v\n", fruits)
	fmt.Printf("veggies: %v\n", veggies)
	fmt.Printf("candy: %v\n", candy)

	allFoods := foods[:] // copy of `foods` array

	fmt.Printf("\nfoods type: %v\n", reflect.TypeOf(foods).Kind())     // —> `array`
	fmt.Printf("allFoods type: %v\n", reflect.TypeOf(allFoods).Kind()) // —> `slice`
}
