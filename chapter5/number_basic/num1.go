package main

import (
	"fmt"
	"strconv"
)

func main() {

	// var a uint = 10 unsigned
	// var b int = -10
	// var c int = 0631
	// var d int = 0x32fa2c75

	// fmt.Println("십진수:", a)
	// fmt.Println("양수:", b)
	// fmt.Println("8진수:", c)
	// fmt.Println("16진수:", d)

	// var a byte = 72
	// var b byte = 0110
	// var c byte = 0x48
	// fmt.Println("십진수:", a)
	// fmt.Println("8진수:", b)
	// fmt.Println("16진수:", c)

	// var asc1 rune = 50556
	// var asc2 rune = 0142574
	// var asc3 rune = 0xc57c
	// fmt.Println("unicode 십진수:", asc1)
	// fmt.Println("unicode 8진수:", asc2)
	// fmt.Println("unicode 16진수:", asc3)

	// fmt.Printf("%c %c %c\n", a, b, c)
	// fmt.Printf("%d %d %d\n", a, b, c)
	// fmt.Printf("%d %o %x\n", a, b, c)

	// fmt.Println("----------------")
	// fmt.Printf("%c %c %c\n", asc1, asc2, asc3)
	// fmt.Printf("%d %d %d\n", asc1, asc2, asc3)
	// fmt.Printf("%d %o %x\n", asc1, asc2, asc3)

	var a float32 = 15.15
	var b float64 = .75851

	var num1 float32 = 14e6
	var num2 float64 = .156875e+3
	var num3 float64 = 5.32521e-10
	var num4 float32 = 10.0
	var num5 float32 = -10.0

	fmt.Println("number:", a)
	fmt.Println("number:", b)
	fmt.Println("number:", num1)
	fmt.Println("number:", num2)
	fmt.Println("number:", num3)
	fmt.Println("number:", num4)
	fmt.Println("number:", num4-0.1)
	fmt.Println("number:", num5-0.1)
	fmt.Println("number:", float32(num4-0.1))
	fmt.Println("number:", float64(num4-0.1))

	fmt.Println("Default int data type mac's CPU: ", strconv.IntSize)

}
