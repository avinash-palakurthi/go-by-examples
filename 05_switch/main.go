package main

import "fmt"

func main() {

	day := "Saturday"

	switch day {
	case "Monday":

		fmt.Println("its Monday ")
	case "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("its weekdays ")

	case "Saturday":
		fmt.Println("its weekend ")

	default:
		fmt.Println("Invalid Day")
	}
}
