package main

import (
	"fmt"
	"log"
)

func main() {

	a, err := validateCapacity(8)

	if err != nil {
		fmt.Println("value: ", a)
		log.Fatal(err)
		//str := err.(notValidNumber).message
		//numb := err.(notValidNumber).numb
		//fmt.Println("message: ", str, numb)
	}

	fmt.Println("validation passed | value:", a)

}

func validateCapacity(i int) (int, error) {

	if i >= 10 {
		return i, nil
	}

	return i, notValidNumber{
		"Not valid number| less than 10| value:",
		i,
	}
}

type notValidNumber struct {
	message string
	numb    int
}

func (ex notValidNumber) Error() string {

	if ex.message != "" {
		return fmt.Sprintf("%v %v:", ex.message, ex.numb)
	}
	return fmt.Sprintf("Not valid number| greater than 10: %v", ex.numb)
}
