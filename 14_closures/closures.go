package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())

	inc1 := counter()
	inc2 := counter()

	fmt.Println(inc1()) // 1
	fmt.Println(inc2()) // 1 (has its own separate state)
}
