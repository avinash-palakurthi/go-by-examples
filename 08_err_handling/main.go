package main

import "fmt"

func division(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("denominator should not be zero")

	}

	return a / b, nil

}

func main() {
	fmt.Println("Error Handling....")
	ans, err := division(105, 2)
	if err != nil {
		fmt.Println("there is an error ", ans)
	}
	fmt.Println(ans)
}
