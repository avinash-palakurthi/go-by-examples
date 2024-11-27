package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Example 1
	number := 80

	if number%2 == 0 {
		fmt.Println(number, ": number is even")
	} else {
		fmt.Println(number, ": number is odd")

	}

	//Example 2

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age: ")

	ageStr, _ := reader.ReadString('\n')
	// Remove trailing newline character
	ageStr = strings.TrimSpace(ageStr)

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid age.")
		return
	}

	if age > 21 {
		fmt.Println("You are allowed...")
	} else {
		fmt.Println("You are not allowed...")
	}

}
