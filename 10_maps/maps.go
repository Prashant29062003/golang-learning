package main

import (
	"fmt"
	"maps"
)

// maps --> hash, object, dict
func main() {
	// creating map

	m := make(map[string]string)

	// settign up the element
	m["name"] = "golang"
	m["area"] = "backend"

	// get the element
	// fmt.Println(m["name"], m["area"])

	// IMP: if no key is defined in map then `zero` value is returns
	// fmt.Println(m["phone"]) // --> gives zero value (here it is string so gives empty string) (if it is integrer it gives 0) (if bool then returns false)

	my := make(map[string]int)

	my["age"] = 23
	my["my_phone"] = 1234567890
	fmt.Println(my["age"])
	fmt.Println(my["phone"]) // --> gives zero value (here it is int so gives 0) (if it is string it gives empty) (if bool then returns false)

	// length of map
	fmt.Println(len(my))

	// print map
	fmt.Println(my) // --> map[age:23 my_phone:1234567890]

	// delete key
	delete(my, "my_phone")

	// print map
	fmt.Println(my) // --> map[age:23]

	// clear the map
	clear(my)

	// print map
	fmt.Println(my) // --> map[]

	// other way to assign map
	m2 := map[string]int{"price": 40, "phones": 9}
	
	v, ok := m2["price"]
	fmt.Println(v)
	if ok {
		fmt.Println("All OK")
		} else {
			fmt.Println("Not OK")
		}
		
		m3 := map[string]int{"price": 40, "phones": 9}

		fmt.Println(maps.Equal(m2,m3))
	}
	