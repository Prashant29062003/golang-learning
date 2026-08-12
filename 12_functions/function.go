package main

import "fmt"

// func add(a int, b int) int {
func add(a, b int) int {	// Parameter shorthand: (a, b int) instead of (a int, b int)
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "javascript", "C++"
}

func getData() (string, string, bool) {
	return "golang", "javascript", true
}

func processIt(fn func(a int) int) int{
	return fn(1)
}

func processIt2 () func (a int) int {
	return func (a int) int {
		return 4
	}
}

func main() {

	result := add(4, 9)
	fmt.Println(result) // 13

	l1, l2, l3 := getLanguages()
	fmt.Println(l1, l2, l3) // golang javascript C++

	d1, _, d3 := getData()
	fmt.Println(d1, d3) // golang true

	fn := func(a int) int {
		return 2
	}

	fmt.Println(processIt(fn))

	fn2 := processIt2()
	fmt.Println(fn2(6)) // return 4
}