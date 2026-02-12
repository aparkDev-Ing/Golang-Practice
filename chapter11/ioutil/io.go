package main

import (
	"fmt"
	"os"
)

func main() {

	s := "hello world"

	err := os.WriteFile("files/hello.txt", []byte(s), 0644)
	errorCheck(err)

	data, err := os.ReadFile("files/hello.txt")

	fmt.Println("Data:", string(data))
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
