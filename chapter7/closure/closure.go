package main

import (
	"fmt"
)

var x1, y1 int = 1, 2

func main() {

	a := func(x int) int {
		x += 1
		return x
	}

	r1 := a(5)
	fmt.Println("result:", r1)

	fmt.Println("-----------------------------------")

	x1, y1 := 5, 10

	b := func(x int) int {
		return x + x1 + y1
	}

	fmt.Println(b(10))

	fmt.Println("-----------------------------------")

	f1 := inc()
	// x := f1()

	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	// fmt.Println(x)
	// fmt.Println(x)
	// fmt.Println(x)

	fmt.Println(inc()())
	fmt.Println(inc()())
	fmt.Println(inc()())

}

func inc() func() int {

	n := 0

	return func() int {
		n++
		return n
	}
}
