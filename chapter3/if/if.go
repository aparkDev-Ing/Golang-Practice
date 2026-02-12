package main

import "fmt"

var i int = 32

func main() {

	var i int = 64

	fmt.Println("Global i: ", i)

	fmt.Println("local i: ", i)

	var a int = 20

	b := 30

	if a >= b {
		fmt.Println("Value a: ", a, " is greater than value b: ", b)
	} else {
		fmt.Println("Value b: ", b, " is greater than value a: ", a)
	}

	if i := 1; i < 5 {
		fmt.Println("C:", i, " is greater than 5")
	}

	if i := 5; i > 10 {
		fmt.Println("value of i is greater than 10", i)
	} else if i := 15; i == 15 {
		fmt.Println("value of i is equal to 15", i)
	}

	var grade int = 85

	if grade > 90 {
		fmt.Println("grade is greater than 90| Grade->", grade)
	} else if grade > 80 && grade < 90 {
		fmt.Println("grade is between 80 and 90| Grade->", grade)
	} else if grade > 70 && grade < 80 {
		fmt.Println("grade is between 70 and 80| Grade->", grade)
	} else if grade > 60 && grade < 70 {
		fmt.Println("grade is between 60 and 70| Grade->", grade)
	} else {
		fmt.Println("grade is less than 60| Grade->", grade)
	}

}
