package main

import (
	"fmt"
)

func rptc(n *int) {
	*n = 77
}
func vptc(n int) {
	n = 77
}

// go 에서는 array가 값타입이라는 증명
func arrayModify1(n [5]int) {
	n[0] = 10
	n[1] = 100
	n[2] = 100
	n[3] = 1000
	n[4] = 10000
}

// go 에서는 pointer를 쓰는방법
func arrayModify2(n *[5]int) {
	n[0] = 10
	n[1] = 100
	n[2] = 100
	n[3] = 1000
	n[4] = 10000

}

func sliceAppend1(n []int) []int {
	//fmt.Printf("memory: %p\n", &n)
	n = append(n, 1, 2, 3)
	fmt.Println("Length of n:", len(n), "cap of n:", cap(n))
	//fmt.Printf("memory address of backing array: %p\n", &n[0])
	return n
}

func sliceModify1(n []int) []int {
	//fmt.Printf("memory: %p\n", &n)
	n[0] = 1
	n[1] = 2
	n[2] = 3
	//fmt.Println("Length of n:", len(n), "cap of n:", cap(n))
	return n
}
func main() {

	var a *int
	var b *int = new(int)

	fmt.Println(a)
	fmt.Println(b)

	i := 7
	fmt.Println(i, &i)

	a = &i
	b = &i

	fmt.Println("Value of a: ", a, "Memory addresss of i: ", &i)
	fmt.Println("Value inside this memory address: ", *a)
	fmt.Println("Memory addresss of a: ", &a)

	fmt.Println("----------------------------------------")

	fmt.Println("Value of b: ", b, "Memory addresss of i: ", &i)
	fmt.Println("Value inside this memory address: ", *b)
	fmt.Println("Memory addresss of b: ", &b)

	var c *int = &i
	d := &i

	//i = 8
	//or
	*d = 8

	fmt.Println("----------------------------------------")

	fmt.Println("Value of c: ", c, "Memory addresss of i: ", &i)
	fmt.Println("Value inside this memory address: ", *c)
	fmt.Println("Memory addresss of c: ", &c)

	fmt.Println("----------------------------------------")

	fmt.Println("Value of d: ", d, "Memory addresss of i: ", &i)
	fmt.Println("Value inside this memory address: ", *d)
	fmt.Println("Memory addresss of d: ", &d)

	fmt.Println("----------------------------------------")

	vi := 7
	//var ia *int = &vi
	ia := &vi

	fmt.Println(vi, *ia, &vi, ia, &ia)

	//역참조 값을올린다
	*ia++
	fmt.Println(vi, *ia)

	fmt.Println("----------------------------------------")

	var ex1 int = 5
	var ex2 int = 5

	vptc(ex1)
	rptc(&ex2)

	fmt.Println(ex1, ex2)

	fmt.Println("----------------------------------------")

	varArr := [5]int{1, 2, 3, 4, 5}
	arrayModify1(varArr)
	fmt.Println(varArr)

	fmt.Println("----------------------------------------")

	varSli1 := make([]int, 5, 5)

	//fmt.Printf("memory:%p\n", &varSli1)
	sliceModify1(varSli1)                     //- even after this backing array itself is same but len still zero
	fmt.Println("slice modify", varSli1[0:3]) //one way to still view using cap for that backing array
	//fmt.Printf("memory: %p\n", &varSli)

	fmt.Println("----------------------------------------")

	varSli := make([]int, 0, 5)

	fmt.Println(varSli)
	sliceAppend1(varSli1) //- even after this backing array itself is same but len still zero
	fmt.Println(varSli)
	fmt.Println("slice append", varSli1[0:3]) //one way to still view using cap for that backing array
	//fmt.Printf("memory: %p\n", &varSli)

	fmt.Println("----------------------------------------")

	varSli2 := make([]int, 5, 5)

	sliceAppend1(varSli2) //
	//varSli2 = sliceAppend((varSli2)) // right way to do it
	fmt.Printf("memory address of backing array: %p\n", &varSli2[0])
	fmt.Println(varSli2[0:5]) //even reslcing won't work because appended 3 more and len becamse 8 cap became 10 which means it re allocated backing array

	fmt.Println("----------------------------------------")

	varArr3 := [5]int{1, 2, 3, 4, 5}
	arrayModify2(&varArr3)
	fmt.Println(varArr3)

}
