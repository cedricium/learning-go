package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{FirstName: firstName, LastName: lastName, Age: age}
}

func main() {
	var people []Person
	for i := 0; i < 10_000_000; i++ {
		people = append(people, MakePerson("alice", "smith", 10))
	}
}
