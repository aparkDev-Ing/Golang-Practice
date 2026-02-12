package main

import "fmt"

//var global int=32

func main() {
	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi! Golang!"
	var k = 4.74
	var l = "Hi! Seoul!"
	var m bool = true

	//a = 5
	//b = "hello world"

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)
	fmt.Println("i: ", i)
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)
	fmt.Println("l: ", l)
	fmt.Println("m: ", m)

	var (
		name      string  = "machine"
		height    int32   = 168
		weight    float32 = 145.0
		isRunning bool    = false
	)

	height = 170
	weight = 135
	isRunning = true
	fmt.Println("name:", name, "height:", height, "weight:", weight, "isRunning:", isRunning)

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	fmt.Println("ShortVa1:", shortVar1, "ShortVa2:", shortVar2, "ShortVa3:", shortVar3)
	shortVar1 = 5
	fmt.Println("After reassign-> ShortVa1:", shortVar1)

	//shortVar1 := 5

	shortVar4, shortVar5 := 5, "hello world"

	fmt.Println("ShortVar4:", shortVar4, "ShortVar5:", shortVar5)

	if i := 10; i < 11 {
		fmt.Println("I: ", i, "is less than 11")
	}

	fmt.Println("i: ", i)

}
