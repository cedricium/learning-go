package main

import "fmt"

func average(nums []float64) (avg float64) {
	total := 0.0

	switch len(nums) {
	case 0:
		avg = 0
	default:
		for _, num := range nums {
			total += num
		}
		avg = total / float64(len(nums))
	}

	return
}

func main() {
	floats := []float64{2.15, 3.14, 42.0, 112.211}
	fmt.Println(average(floats))
}
