package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("http request")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		fmt.Println("Error while fetching data ", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("type of response : ", res)

	//todo :- read res from the body
	data, err := io.ReadAll(res.Body)
	if err != nil {

		fmt.Println("Error whlie reading data", err)
		return
	}
	fmt.Println("response : ", string(data))

}
