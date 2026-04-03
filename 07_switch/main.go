package main

import (
	"fmt"
	"time"
)

func main() {
	// Sample switch
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Other")
	}


	// Multiple conditional switch
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's workday.")
	}

	whoAmI := func (i interface{})  {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer.")
		case string:
			fmt.Println("It's a string.")
		case float32:
			fmt.Println("It's a ratiotional value.")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Printf("Other type %T.\n", t)
		}
	}

	whoAmI(34e2) // Other type float64
	whoAmI("hey")
}
