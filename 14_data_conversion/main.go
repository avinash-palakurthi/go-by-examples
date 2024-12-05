package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var num int = 12

	// fmt.Printf(" num type is  %T\n", num)

	// var data float64 = float64(num)
	// fmt.Printf("data type of Data: %T\n", data)

	numString := "1234"
	fmt.Printf("numString = %s, numString data type : %T\n", numString, numString)

	stringNum, _ := strconv.Atoi(numString)
	fmt.Printf("stringNum = %d, stringNum data type :%T\n", stringNum, stringNum)

}
