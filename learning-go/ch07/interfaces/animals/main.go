package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "woof"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "meow"
}

type Gopher struct{}

func (g Gopher) Speak() string {
	return "squeak"
}

type Programmer struct{}

func (p Programmer) Speak() string {
	return "leetcode"
}

type Sloth struct{}

func main() {
	// animals := []Animal{Dog{}, Cat{}, Gopher{}, Programmer{}, Sloth{}}
	//                                                           ~~~~~~~
	// 	cannot use Sloth{} (value of type Sloth) as Animal value in array or slice
	// 	literal: Sloth does not implement Animal (missing method Speak)

	animals := []Animal{Dog{}, Cat{}, Gopher{}, Programmer{}}
	for _, a := range animals {
		fmt.Println(a.Speak())
	}
}
