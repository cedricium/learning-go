package main

import "fmt"

func main() {
	/* note: using an ellipsis allows the compiler to determine the size of the array */
	deploymentOptions := [...]string{"Heroku", "AWS", "GCP", "Azure"}

	fmt.Println(len(deploymentOptions))

	for index, option := range deploymentOptions {
		fmt.Printf("Index: %v	Item: %v\n", index, option)
	}
}
