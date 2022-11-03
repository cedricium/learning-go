package main

import "fmt"

func main() {
	var deploymentOptions = [4]string{"Heroku", "AWS", "GCP", "Azure"}

	for i := 0; i < len(deploymentOptions); i++ {
		option := deploymentOptions[i]
		fmt.Printf("Index: %v	Item: %v\n", i, option)
	}
}
