package main

import (
	"fmt"
	"log"
)

func validate(i int) (int, error) {

	if i < 10 {
		return i, fmt.Errorf("Value is less than 10| input value: %v", i)
		//return i, errors.New("Value is less than 10| ")
	}

	return i, nil
}

func main() {

	value, err := validate(8)

	if err != nil {
		//fmt.Println("Input Value: ", value)
		log.Fatal(err)
	}

	fmt.Println("value is:", value)
}
