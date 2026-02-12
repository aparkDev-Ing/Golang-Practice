package main

//when program starts, this initalize happenes even before main method.

import (
	"fmt"

	checkSum "github.com/aparkDev-Ing/Golang-Practice/chapter4/lib"
)

func init() {

	fmt.Println("Init1: Init1")

}

func init() {

	fmt.Println("Init2: Init2")

}

func init() {

	fmt.Println("Init3: Init3")

}

func main() {

	fmt.Println("Main")
	fmt.Println(checkSum.CheckNum(100))
}
