package main

import "fmt"

func UpdateSlice(ss []string, s string) {
	ss[len(ss)-1] = s
}

func GrowSlice(ss []string, s string) {
	ss = append(ss, s)
	fmt.Println(ss) // 3	[hello world universe galaxy]
}

func main() {
	ss := []string{"hello", "world", "goodbye"}

	fmt.Println(ss) // 1	[hello world goodbye]
	UpdateSlice(ss, "universe")
	fmt.Println(ss) // 2	[hello world universe]
	GrowSlice(ss, "galaxy")
	fmt.Println(ss) // 4	[hello world universe]
}
