package main

import (
	"fmt"
	"strconv"
)

func PrintStatement(m string) {
	fmt.Println("Hello:", m)
}

func calculate1(x int, y int) int {
	return x + y
}

func calculateTotal(pretax int, f func(int) float64) float64 {

	return f(pretax)
}

func calculateTax(pretax int) float64 {

	return float64(pretax) * 0.9
}

func incValues(varArray *[5]int) {

	for i := 0; i < len(varArray); i++ {
		varArray[i]++
	}

}

func multiplyAndAdd(x int, y int) (int, int) {
	return x * y, x + y
}

func divide(x int, y int) (r1 int, r2 int) {
	r1 = x
	r2 = y
	return r1 / r2, r2 / r1
}

func divide2(x int, y int) (int, int) {
	r1 := x
	r2 := y
	return r1 / r2, r2 / r1
}

func multiply(n ...int) int {
	val := 1
	for _, value := range n {
		val *= value
	}

	return val
}

func sum(n int, y int) (int, int) {

	return n + 1, y + 1
}
func minus(n int, y int) (int, int) {

	return n - 1, y - 1
}

func fact(n int) int {

	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	PrintStatement("Aaron")
	fmt.Println(calculate1(5, 10))

	fmt.Println(strconv.Itoa(calculate1(5, 10)))

	fmt.Println("-------------------------------------------")

	pretaxTotal := 100

	fmt.Println("Total Pay after tax:", calculateTotal(pretaxTotal, calculateTax))

	fmt.Println("-------------------------------------------")

	varArray := [5]int{1, 2, 3, 4, 5}
	incValues(&varArray)
	fmt.Println(varArray)

	fmt.Println("-------------------------------------------")

	a, b := multiplyAndAdd(5, 10)
	//c,_ :=multiplyAndAdd(5,10)
	//_, d := multiplyAndAdd(5, 10)
	fmt.Println("multiplication:", a, "addition:", b)

	fmt.Println("-------------------------------------------")
	r1, r2 := divide(100, 5)
	fmt.Println(r1, r2)

	fmt.Println("-------------------------------------------")

	fmt.Println(multiply(4, 10, 2))

	aArr := []int{1, 2, 3}
	fmt.Println(multiply(aArr...))

	fmt.Println("-------------------------------------------")

	f1 := []func(int, int) (int, int){sum, minus}
	//f1 = append(f1, sum)

	sum1, sum2 := f1[0](1, 2)
	minus1, minus2 := f1[1](1, 2)
	fmt.Println("sum1:", sum1, "sum2:", sum2)
	fmt.Println("minus1:", minus1, "minus2:", minus2)

	fmt.Println("-------------------------------------------")
	var f2 func(int, int) (int, int) = sum

	fmt.Println(f2(5, 10))
	fmt.Println("-------------------------------------------")

	m := map[string]func(int, int) (int, int){
		"function1": sum,
		"function2": minus,
	}

	fmt.Println(m["function1"](10, 5))
	fmt.Println(m["function2"](10, 5))

	fmt.Println("-------------------------------------------")

	fmt.Println(fact(5))

	fmt.Println("-------------------------------------------")

	func() {
		fmt.Println("Annoymous funciton")
	}()

	fmt.Println("-------------------------------------------")

	func(m string) {
		fmt.Println("ex2: ", m)
	}("hello")

	fmt.Println(func(x int, y int) int {
		return x * y
	}(5, 10))

	// f3 := func(x int, y int) int {
	// 	return x * y
	// }
	// fmt.Println(f3(5, 10))

	// f3 := func(x int, y int) int {
	// 	return x * y
	// }(5, 10)

	// fmt.Println(f3)

	fmt.Println(calculateTotal(10, func(x int) float64 {

		return float64(x) * 0.9
	}))
}
