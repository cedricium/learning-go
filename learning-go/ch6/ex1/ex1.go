package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{FirstName: firstName, LastName: lastName, Age: age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{FirstName: firstName, LastName: lastName, Age: age}
}

func main() {
	p1 := MakePerson("Cedric", "Diggory", 28)
	fmt.Println(p1)
	p2 := MakePersonPointer("Harry", "Potter", 25)
	fmt.Println(p2)
}

/*
$ go build -gcflags="-m" ex1.go
# command-line-arguments
./ex1.go:11:6: can inline MakePerson
./ex1.go:15:6: can inline MakePersonPointer
./ex1.go:20:18: inlining call to MakePerson
./ex1.go:21:13: inlining call to fmt.Println
./ex1.go:22:25: inlining call to MakePersonPointer
./ex1.go:23:13: inlining call to fmt.Println
./ex1.go:11:17: leaking param: firstName to result ~r0 level=0
./ex1.go:11:28: leaking param: lastName to result ~r0 level=0
./ex1.go:15:24: leaking param: firstName
./ex1.go:15:35: leaking param: lastName
./ex1.go:16:9: &Person{...} escapes to heap
./ex1.go:21:13: ... argument does not escape
./ex1.go:21:14: p1 escapes to heap
./ex1.go:22:25: &Person{...} escapes to heap
./ex1.go:23:13: ... argument does not escape
*/
