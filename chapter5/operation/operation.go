package main

import (
	"fmt"
	"math"
)

func main() {

	// //alias for uint8 type
	// var a byte = math.MaxUint8

	// //alias for int32 type
	// var b rune = math.MaxInt32

	// var c int16 = math.MaxInt16

	// var d int32 = math.MaxInt32

	// var e int64 = math.MaxInt64

	// fmt.Println(a, b, c, d, e, math.MaxFloat32, math.MaxFloat64)

	// n1 := 10000
	// n2 := int16(10002)
	// n3 := uint8(100)
	// n4 := uint8(255)

	// fmt.Println(uint8(n2) > n3)

	// fmt.Println(int16(n1) + n2)
	// //fmt.Println(n1 + int(n2))
	// fmt.Println(n2 > int16(n3))
	// fmt.Println(n4 + 1)

	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n2 - n1)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println(n1 << 2)
	fmt.Println(n1 >> 2)
	fmt.Println(^n1)

	var n3 uint16 = 1500
	var n4 float32 = 15.5
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	fmt.Println(float32(n3) + n4)
	fmt.Println(n3 + uint16(n4))
	fmt.Println(n5 + uint16(n6))

	var n7 uint8 = math.MaxUint8
	fmt.Println("overflow| 256 % 256", n7+1)
	//fmt.Println(256 % 256)

	//var n8 uint32 = -1

}
