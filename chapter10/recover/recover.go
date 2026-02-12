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
		if s := recover(); s != nil {
			fmt.Println("Error message:", s)
		}
	}()

	a := [3]int{1, 2, 3}

	for i := 0; i < 4; i++ {
		fmt.Println("value: ", a[i])
	}
}

func main() {

	fmt.Println("main starts")
	test2()
	fmt.Println("main ends")
}
