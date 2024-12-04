package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 40
}

func main() {

	//! long method
	// var num string
	// num = "3"
	// var ptr *string
	// ptr = &num

	//! short method

	num := 2
	ptr := &num

	fmt.Println("num : ", ptr)
	fmt.Println("ptr : ", *ptr)

	value := 25

	modifyValueByReference(&value)
	fmt.Println("value is :", value)

}
