package main

import (
	"errors"
	"fmt"
)

func main() {

	err1 := errors.New("Error occurred -1")

	fmt.Println("Error1:", err1)
	fmt.Println("Error1:", err1.Error())

}
