package main

import "fmt"

// Slices => array of dynamic length
// most used construct in go
// + usefull methods

func main() {
	// un-initilized slice is by default nil
	// var nums []int

	// fmt.Println(nums) // []
	// fmt.Println(nums == nil) // True
	// fmt.Println(len(nums)) // 0

	var nums = make([]int, 2, 5)
	// Capacity -> maximum numbers of elements can fit but it automatically update the size as it is dynamic
	fmt.Println(cap(nums)) // 2

	fmt.Println(nums) // [0,0]
	// fmt.Println(nums == nil) // False as it has two elements initially

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	fmt.Println(nums)
	fmt.Println(len(nums)) // here size is automatilly increase to 10 from 5 as it exceeds number 5 it's size increases double
}
