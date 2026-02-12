package main

import (
	"fmt"
	"math"
)

func Power(f float64, i float64) (float64, error) {

	if math.IsNaN(f) {
		return 0, powError{
			"Input is invalid. Please enter float64 type",
			f,
		}
	}

	if math.IsNaN(i) {
		return 0, powError{
			"Input is invalid. Please enter float64 type",
			f,
		}
	}

	// if math.IsNaN(f) {
	// 	return 0, powError{
	// 		"Input is invalid. Please enter float64 type",
	// 		f,
	// 	}
	// }

	// if math.IsNaN(i) {
	// 	return 0, powError{
	// 		"Input is invalid. Please enter float64 type",
	// 		f,
	// 	}
	// }

	if f == 0 {
		//return 0, errors.New("cannot use 0")
		//return 0, fmt.Errorf("cannot use 0")
		//custom type
		return 0, powError{
			"Input cannot be 0",
			f,
		}
	}

	return math.Pow(f, i), nil
}

func main() {

	a, err := Power(3, 3)

	if err != nil {
		fmt.Println("error occured:", err)
	} else {
		fmt.Println("value:", a)
	}
}

type powError struct {
	message string
	number  float64
}

func (ex powError) Error() string {
	if ex.message != "" {
		return fmt.Sprintf("Input value: %v | Message: %v", ex.number, ex.message)
	}

	return fmt.Sprintf("Input value: %v | Message: %v", ex.number, "Exception Occured!")
}

// package main

// import (
// 	"fmt"
// 	//"math"
// )

// func Power(f any, i interface{}) (float64, error) {

// 	f, ok1 := f.(float64)

// 	i, ok2 := i.(float64)

// 	if !ok1 {
// 		return 0, powError{
// 			"Input is invalid. Please enter float64 type",
// 			f,
// 		}
// 	}

// 	if !ok2 {
// 		return 0, powError{
// 			"Input is invalid. Please enter float64 type",
// 			f,
// 		}
// 	}

// 	// if math.IsNaN(f) {
// 	// 	return 0, powError{
// 	// 		"Input is invalid. Please enter float64 type",
// 	// 		f,
// 	// 	}
// 	// }

// 	// if math.IsNaN(i) {
// 	// 	return 0, powError{
// 	// 		"Input is invalid. Please enter float64 type",
// 	// 		f,
// 	// 	}
// 	// }

// 	if f == 0 {
// 		//return 0, errors.New("cannot use 0")
// 		//return 0, fmt.Errorf("cannot use 0")
// 		//custom type
// 		return 0, powError{
// 			"Input cannot be 0",
// 			f,
// 		}
// 	}

// 	return 0, nil
// }

// func main() {

// 	a, err := Power("hello", 2)

// 	if err != nil {
// 		fmt.Println("error occured:", err)
// 	} else {
// 		fmt.Println("value:", a)
// 	}
// }

// type powError struct {
// 	message any
// 	number  any
// }

// func (ex powError) Error() string {
// 	if ex.message != "" {
// 		return fmt.Sprintf("Input value: %v | Message: %v", ex.number, ex.message)
// 	}

// 	return fmt.Sprintf("Input value: %v | Message: %v", ex.number, "Exception Occured!")
// }
