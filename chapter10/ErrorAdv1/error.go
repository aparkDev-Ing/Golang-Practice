package main

import (
	"errors"
	"fmt"
	"math"
)

func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("cannot use 0")
		//return 0, fmt.Errorf("cannot use 0")
	}

	return math.Pow(f, i), nil
}

func main() {

	a, err := Power(2, 2)

	if err != nil {
		fmt.Println("error occured:", err)
	} else {
		fmt.Println("value:", a)
	}
}
