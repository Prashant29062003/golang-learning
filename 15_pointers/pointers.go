package main

import "fmt"

// Passing a standard value
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}


// Passing a pointer value
func changeNumRef(num *int){
	*num = 4
	fmt.Println("In changeNumRef", *num)
}


// Technical Correction: "Pass by Value" vs. "Pass by Reference"
// The Misconception Callout:

// Languages like C++ or C# feature true reference variables (e.g., void change(int &num)), where the function argument becomes an alias for the exact same variable in the caller scope without using explicit memory pointers. Go does not have this feature.

// How Go Handles Pointers:

// When you call changeNumRef(&num), Go creates a copy of the memory address (*int). Because this copied pointer still points to num's address in memory, dereferencing it (*num = 4) mutates the original variable's underlying value.

func main() {
	num := 1

	changeNum(num)

	fmt.Println("After changeNum in main", num)
	
	fmt.Println("--------------------------------------------")
	
	// fmt.Println(&num)
	changeNumRef(&num)
	fmt.Println("After changeNumRef in main", num)
}