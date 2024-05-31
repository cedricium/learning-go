package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	ss1 := greetings[:2]
	ss2 := greetings[1:4]
	ss3 := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(ss1)
	fmt.Println(ss2)
	fmt.Println(ss3)
}
