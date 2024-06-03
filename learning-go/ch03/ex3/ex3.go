package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	// struct literal, no keys
	alice := Employee{"alice", "baker", 1}

	// struct literal, using keys
	bob := Employee{
		firstName: "bob",
		lastName:  "smith",
		id:        2,
	}

	// var declaration, dot notation
	var carol Employee
	carol.firstName = "carol"
	carol.lastName = "nguyen"
	carol.id = 3

	fmt.Println(alice)
	fmt.Println(bob)
	fmt.Println(carol)
}
