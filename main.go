package main

import (
	"avinash-palakurthi/go-by-examples.git/car"
	"avinash-palakurthi/go-by-examples.git/greet"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("this main executable function")

	// folderName + functionName
	greet.PrintGreetings("good morning ")
	car.MyCar("Bmw", "Germany")

	fmt.Println("Enter Your Name")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println("good morning", name)

}
