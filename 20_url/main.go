package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("1 . This is url example...")

	myUrl := "https://example.com:8080/path/to/resourse?key=value1&key=value2"

	parsedURL, err := url.Parse(myUrl)

	if err != nil {

		fmt.Println("cant parse url", err)
		return
	}

	fmt.Printf("Type of url : %T \n", parsedURL)
	fmt.Println("Scheme :", parsedURL.Scheme)
	fmt.Println("Host :", parsedURL.Host)
	fmt.Println("Hostname :", parsedURL.Hostname())
	fmt.Println("Path :", parsedURL.Path)
	fmt.Println("RawQuery :", parsedURL.RawQuery)

	//todo : modifying url values from parsed url boject

	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=JOHN"

	//todo : converting parsed url back to string
	newUrl := parsedURL.String()
	fmt.Println("new url :", newUrl)
}
