package main

import (
	"fmt"

	arithmetic "github.com/aparkDev-Ing/Golang-Practice/chapter12/packages/arithmetic"
)

func main() {

	fmt.Println("hello world")
	ar1 := arithmetic.Numbers{
		10,
		10,
	}

	minused := ar1.Minus()

	fmt.Println(minused)

}
