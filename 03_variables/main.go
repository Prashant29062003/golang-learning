package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 3
	fmt.Println(b + c)
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)
	f := "apple"
	fmt.Println(f)

	var g float64
	fmt.Println(g)

	fmt.Println(a == f)

	/*
	var username string = "Prashant"
	fmt.Println(username)

	// Types 
	fmt.Printf("Variables is of type %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)

	// Types 
	fmt.Printf("Variables is of type %T\n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type %T\n", smallVal)

	var smallFloat32 float32 = 32.438324103948103488109483809483048380434901311111111111111111118
	fmt.Println(smallFloat32) // 32.438324
	fmt.Printf("Variables is of type %T\n", smallFloat32)

	var smallFloat64 float64 = 32.438324103948103488109483809483048380434901311111111111111111118
	fmt.Println(smallFloat64) // 32.438324103948105
	fmt.Printf("Variables is of type %T\n", smallFloat64)
	*/

	var intVar int
	fmt.Println(intVar)
	fmt.Printf("Type of this is %T\n", intVar)

	var stringVar string
	fmt.Println(stringVar)
	fmt.Printf("Type of this is %T\n", stringVar)

	var boolVar bool
	fmt.Println(boolVar)
	fmt.Printf("Type of this is %T\n", boolVar)


	// implicit type

	var web = "Prashant.in"
	fmt.Println(web)


	numberOfUser := 30002.54
	fmt.Println(numberOfUser)
	fmt.Printf("Type of %.2f is %T\n", numberOfUser, numberOfUser) // %f ==> for specific float value
	fmt.Printf("Type of %v is %T\n", numberOfUser, numberOfUser) // %v ==> is for general value
}
