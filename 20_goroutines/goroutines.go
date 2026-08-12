package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing task: ",id)
}

func main() {
	var wg sync.WaitGroup // 1. Create a WaitGroup

	for i := 0; i < 1000; i++{
		wg.Add(1) // 2. Increment the counter BEFORE launching the goroutine
		go task(i, &wg)

		// go func (i int) {
		// 	defer wg.Done()  // 3. Decrement the counter when this function finishes
		// 	fmt.Println(i) 
		// } (i)
	}

	wg.Wait() // 4. Block here until the counter reaches 0
	fmt.Println("All tasks completed!")
}







