package main

import "fmt"

func main() {
	//! method 1
	studentGrades := make(map[string]int)

	studentGrades["shiva"] = 498
	studentGrades["venku"] = 199
	studentGrades["seetha"] = 299
	studentGrades["hanu"] = 990
	studentGrades["hari"] = 109

	fmt.Println("shiva marks :", studentGrades["shiva"])
	fmt.Println("venku marks :", studentGrades["venku"])

	// checking student exists

	grade, exist := studentGrades["venku"]

	fmt.Println("venku grade :", grade)
	fmt.Println("does venku exists :", exist)

	//for loop
	for index, value := range studentGrades {
		fmt.Printf(" %s got   marks  %d \n", index, value)
	}

	//! method 2

	person := map[string]int{
		"surya":   21,
		"chandra": 22,
	}

	for index, value := range person {
		fmt.Printf("-------%s got %d \n", index, value)
	}
}
