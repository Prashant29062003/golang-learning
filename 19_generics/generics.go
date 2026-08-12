package main

import (
	"fmt"
)

// duplication for same logic ---
// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// func printString(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// ------------------------------

// Using generics we can use same logic once and apply any where with any type
// func printData[T any](items []T){
// func printData[T int | string](items []T){
func printData[T comparable, V string | int](items []T, phone V){
	for _, item := range items{
		fmt.Println(item, phone)
	}
}


type stack[T any] struct {
	elements []T
}

func (s *stack[T]) Push (item T){
	s.elements = append(s.elements, item)
}



func main() {
	nums1 := []int {1,2,3}
	printData(nums1, "1234567890")

	names1 := []string {"go", "js", "python"}
	printData(names1, 1234567890)

	myStack1 := stack[string]{
		elements: []string{"ved", "rahul"},
	}
	myStack1.Push("rupa")
	
	myStack2 := stack[int]{
		elements: []int{1, 3, 90},
	}
	myStack3 := stack[bool]{
		elements: []bool{true, false, true},
	}

	fmt.Println(myStack1)
	fmt.Println(myStack2)
	fmt.Println(myStack3)
}