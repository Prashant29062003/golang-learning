package main

import "fmt"

func main() {
	// slices
	nums := []int{6, 7, 9}

	// simple for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	// // via range
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println(sum)

	for i, num := range nums{
		fmt.Println(i, num)
		// 0 6
		// 1 7
		// 2 9
	}

	// map
	m := map[string]string{"fname": "John", "lname": "Doe"}

	for k, v := range m{
		fmt.Println(k,v)
		// fname John
		// lname Doe
	} 

	// string --> range
	for i, c := range "golang"{
		fmt.Println(i,c)
		// starting byte of rune, unicode code point rune
		// 0 103
		// 1 111
		// 2 108
		// 3 97
		// 4 110
		// 5 103
	}
}