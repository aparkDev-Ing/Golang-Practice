package main

import "fmt"

func main() {

	//initialization required
	const a string = "final word"
	const b = "Test2"
	const c int32 = 10 * 10

	const e = 35.6

	//const g=getHeight() //not possible because getHeight can return anything

	fmt.Println("a:", a, "b:", b, "c:", c, "e:", e)

	//const e1,f ,g= 1, 2, 3

	const (
		name   string  = "aaron"
		height int16   = 168
		weight float32 = 145
	)

	fmt.Println(name, height, weight)

}
