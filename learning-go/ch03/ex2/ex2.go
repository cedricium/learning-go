package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	runes := []rune(message)      // first, decode `message` string into slice of runes
	fmt.Println(string(runes[3])) // finally, convert and print string representation of rune at index 3
}
