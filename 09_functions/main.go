package main

import "fmt"

func simpleFunction() {
	fmt.Println("this is simple function")
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) (result int) {

	result = a * b
	return
}

func main() {
	fmt.Println("this is main function")
	simpleFunction()
	adding := add(2, 6)
	fmt.Println("sum of two numbers :", adding)

	multi := multiply(2, 3)
	fmt.Println(multi)
}
