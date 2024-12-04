package main

import "fmt"

func main() {
	//! method 1
	// numbers := []int{1, 2, 3}
	// numbers = append(numbers, 20, 30, 40)

	// fmt.Println(numbers)
	// fmt.Println("length :", len(numbers))
	// fmt.Printf("data type of nmubers : %T\n", numbers)

	/*fmt.Println("slice :", numbers)
	fmt.Println("length :", len(numbers))
	fmt.Println("capacity:", cap(numbers))*/

	//! method 2
	//todo:  This method is flexible,this method starts with [ MAKE fn ] .  we have togive length & capacity

	nums := make([]int, 3, 5) // 3= length & 5=capacity
	//!wrong method of appending
	// nums[0] = 1
	// nums[1] = 2
	// nums[2] = 3
	// nums[3] = 4

	//! this is right method

	nums = append(nums, 21, 31, 41, 99, 89, 79, 69, 59)

	fmt.Println("slice :", nums)
	fmt.Println("length :", len(nums))
	fmt.Println("capacity :", cap(nums))

}
