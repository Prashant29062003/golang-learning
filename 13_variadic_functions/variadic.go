package main

import "fmt"

func sum (nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
func main() {
	// function where we can pass n number of parameters 
	fmt.Println(1, 2, 3, 4, 5, "...")
	fmt.Println(sum(1,2,3,4))

	// slice
	nums2 := []int {3, 4, 6, 8}
	fmt.Println(sum(nums2...))
} 