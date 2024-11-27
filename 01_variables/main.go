package main

import "fmt"

func main() {
	/*
			var declares 1 or more variables.
		  You can declare multiple variables at once

	*/
	var first = "hello"
	fmt.Println("first:", first)

	var variable1, variable2 int = 2, 21
	fmt.Printf("variable1 = %d \n, variable2 = %d \n", variable1, variable2)

	// Variables declared without a value  initialization are shown as  zero-valued, type : string shows empty space,

	var empty_valie string
	var boolean_value bool
	var zero_value int
	fmt.Println("zero_value  default value is :", zero_value)
	fmt.Println("empty_value  default value is :", empty_valie)
	fmt.Println("boolean_value  default value is :", boolean_value)

	// shorthand variable declaration

	fruit := "Mango"
	fmt.Println("Fruit : ", fruit)
}
