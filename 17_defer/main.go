package main

import "fmt"

func main() {

	defer fmt.Println("this is 1st")
	defer fmt.Println("this is 2nd")
	fmt.Println("this is 3rd")

}
