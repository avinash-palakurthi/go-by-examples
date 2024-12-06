package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("currentTime : ", currentTime)

	formatted := currentTime.Format("01-02-2006, Monday, 3:04 PM")

	fmt.Println(formatted)
}
