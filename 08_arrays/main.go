package main

import "fmt"

// numbered sequence of specific length
func main(){
	var nums [4]int

	/*
	/ ---  ---
	*/
	fmt.Println(nums) // [0,0,0,0]
	fmt.Println(len(nums)) // 4

	// add number
	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(nums) // [1,0,0,0]

	var arr [4]bool
	arr[2] = true
	fmt.Println(arr)

	var names [4]string
	names[0] = "golang"
	names[1] = "_"
	names[2] = "_"
	names[3] = "_"
	fmt.Println(names)

	/*
	/ ---  ---
	*/

	// to declare in single line
	nums2 := [3]int{1,2,3}
	fmt.Println(nums2)

	// 2D Array
	nums3 := [2][2]int{{3,4},{5,6}}
	fmt.Println(nums3)


	// Benifits
	// - Fixed size, that is predictable
	// - Memory Optimization
	// - Constant Time optimization
	
}