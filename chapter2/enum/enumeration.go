package main

import "fmt"

func main() {

	// const (
	// 	Monday   string = "Monday"
	// 	Tuesday  string = "Tuesday"
	// 	Wendsday string = "Wendsday"
	// 	Thursday string = "Thursday"
	// 	Friday   string = "Friday"
	// )

	//fmt.Println(Monday, Tuesday, Wendsday, Thursday, Friday)

	// const (
	// 	A = iota * 10
	// 	B
	// 	C
	// )

	// fmt.Println(A, B, C)

	// const(
	// 	_=iota
	// 	A
	// 	_
	// 	C
	// 	D
	// )

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		_
		Diamond
	)

	fmt.Println(DEFAULT, SILVER, GOLD, Diamond)

}
