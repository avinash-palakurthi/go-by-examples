package main

import "fmt"

func main() {
	//for loop construct to iterate over a sequence of values

	//Example - 1
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	//Example - 2
	sum := 0

	for sum <= 10 {
		sum += 2
		fmt.Println("SUM 2 : ", sum)
	}

	//Example 3: For Loop as a While Loop

	num := 0
	for {
		if num > 10 {
			break
		}
		fmt.Println("num :", num)
		num++
	}

	// Example 4: Iterating Over an Array

	// declaring variable with specific condition

	//!Array (fixed size of Array)
	//!Declaration: var arrayName [size]type
	//! Example: numbers := [5]int{1, 2, 3, 4, 5}

	// numbers := [3]int{1, 2, 3}

	//! Dynamic size: Slices can grow or shrink dynamically.
	//! Declaration: var sliceName []type

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for index, number := range numbers {
		fmt.Printf("Index: %d, Number: %d\n", index, number)
	}
}
