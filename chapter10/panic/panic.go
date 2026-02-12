package main

import (
	"fmt"
	"log"
)

func test1() {
	defer fmt.Println("Must print")

	log.Panic("error occurred")

}

func test2() {

	defer func() {
		s := recover()
		fmt.Println("Error message:", s)

	}()

	// s := 0

	// v := 5

	// fmt.Println(v / s)
	panic("Error Occurred!")
}

func main() {

	// fmt.Println("main starts")
	// panic("Error Occurred")
	// fmt.Println("main ends")

	fmt.Println("main starts")
	test2()
	fmt.Println("main starts")
}
