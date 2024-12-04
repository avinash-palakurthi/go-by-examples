package main

import "fmt"

func main() {
	//! 1. Creating an Array of Strings:
	/*var fruits [3]string

	fruits[1] = "Mango"

	fmt.Println(fruits)
	fmt.Println(len(fruits))*/

	//! 2. Initializing an Array During Declaration:
	/*

		cars := []string{"skoda", "VW", "Audi", "Benz"}
		fmt.Println(cars)
		fmt.Println(len(cars))
	*/

	// b := []int{10, 20, 30, 40, 50}
	// fmt.Println("dcl:", b)

	// b = []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println("dcl:", b)

	// b = []int{100, 3: 2100, 500}
	// fmt.Println("idx:", b)

	var name [4]string
	name[1] = "avi"
	name[3] = "shiv"

	fmt.Println("total length of name :", len(name))
	fmt.Printf("%q\n", name)
}
