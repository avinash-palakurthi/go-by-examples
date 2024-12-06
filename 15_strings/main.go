package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings ....")
	// fruits := "mango, apple, banana"
	// data := strings.Split(fruits, ",")

	// fmt.Println(data)

	// words := "one one two three four four four "
	// noOfRepeatedWords := strings.Count(words, "five")

	// fmt.Println(noOfRepeatedWords)

	// str := "   hello , ok     "

	// trimmed := strings.TrimSpace(str)
	// fmt.Println("trimmed :", trimmed)

	str1 := "volks"
	str2 := "wagen"

	completeString := strings.Join([]string{str1, str2}, " ")

	fmt.Println(completeString)
}
