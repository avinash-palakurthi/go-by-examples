package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("file handling....")

	//! file creating

	/*
		file, errors := os.Create("example.txt")

		if errors != nil {
			fmt.Println("problem occured while creating file... ", errors)
			return
		}

		defer file.Close()

		content := "hello world"
		byte, err := io.WriteString(file, content+"\n")
		fmt.Println("total bytes :", byte)
		if err != nil {
			fmt.Println("problem occured while creating file... ", err)
			return
		}

		fmt.Println("successfully created file")

	*/

	//! file open/reading

	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("error while opening file", err)
		return
	}

	defer file.Close()

	//todo:- creating buffer to read the file
	buffer := make([]byte, 1024)

	//todo:-  reading the file

	for {
		n, err := file.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("error while reading the file", err)
			return

		}

		//todo :-  process of reading file
		fmt.Println(string(buffer[:n]))
	}

}
