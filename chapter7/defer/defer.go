package main

import (
	"fmt"
)

func f1() {
	fmt.Println("f1 starts")
	defer f2()
	fmt.Println("f2 end")
}

func f2() {
	fmt.Println("f2 invoked")
}

func f3(s string) {
	fmt.Println("f3 starts ")
	fmt.Println("f3 end ")
	defer func(s string) {
		fmt.Println("callback1 function triggers:", s)
	}(s)

	defer func(s string) {
		fmt.Println("callback2 function triggers:", s)
	}(s)
}

func f4(x int) {
	fmt.Println("f4 starts")
	defer fmt.Println(multiply(add(x)))
	fmt.Println("f4 ends")
}

func f5(x string) {
	fmt.Println("f4 starts")
	defer print2(print1(x))
	fmt.Println("f4 ends")
}

func print1(x string) string {

	fmt.Println("print1", x)
	return x
}

func print2(x string) {

	fmt.Println("print2", x)
}

func add(x int) int {

	fmt.Println("add invoked")
	return x + 1
}

func multiply(x int) int {

	fmt.Println("multiply invoked")
	return x * 2
}

func stack(i int) {
	fmt.Println("functin starts")
	for i < 10 {
		defer fmt.Println("value: ", i)
		i++
	}
	fmt.Println("functin ends")
}

func main() {

	f1()

	fmt.Println("------------------------------------")

	f3("callback function!")

	fmt.Println("------------------------------------")
	stack(0)

	fmt.Println("------------------------------------")
	f4(6)
	//f5("hello")
}
