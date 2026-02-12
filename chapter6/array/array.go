package main

import (
	"fmt"
)

func main() {

	var intArr2 [5]int
	var intArr3 [5]int = [5]int{1, 2, 3, 4, 5}
	var intArr4 = [5]int{1, 2, 3, 4, 5}

	fmt.Println(intArr2, intArr3, intArr4)

	intArr5 := [5]int{1, 2, 3, 4, 5}
	intArr6 := [5]int{1, 2, 3}
	intArr7 := [...]int{1, 2, 3}

	fmt.Println(intArr5, intArr6, intArr7)

	intArr8 := [5][2]int{
		{1, 2}, {1, 2}, {1, 2},
	}
	//intArr3:= make([]int,3)

	intArr2[2] = 5
	fmt.Println(intArr2)

	fmt.Println(len(intArr8))

	strArr1 := [5]string{"hello", "hi", "Aaron", "david", "claire"}

	fmt.Println(strArr1)

	fmt.Println("--------------------------------------------------------")

	for i := 0; i < len(strArr1); i++ {
		fmt.Println(strArr1[i])
	}

	for k, v := range strArr1 {
		fmt.Println("index:", k, "value:", v)
	}

	//값만 출력할대
	for _, v := range strArr1 {
		fmt.Println("value:", v)
	}

	//index만 출력할때
	for v, _ := range strArr1 {
		fmt.Println("value:", v)
	}

	//skip시 첫번째, 즉인덱스만 출력
	for v := range strArr1 {
		fmt.Println("value:", v)
	}

	fmt.Println("-----------------------")

	arrInt1 := [5]int{1, 2, 3, 4, 5}
	arrInt2 := arrInt1

	fmt.Printf("%p %v\n", &arrInt1, arrInt1)
	fmt.Printf("%p %v\n", &arrInt2, arrInt2)

	a := []int{1, 2, 3}

	b := a

	c := append(a, 1)

	fmt.Println(c)
	fmt.Println(a, b)

	//fmt.Println(a[:4])

	fmt.Println("--------------------------")

	a1 := make([]int, 0, 2)

	a2 := append(a1, 1, 2)

	fmt.Println(a2)
	fmt.Println(a1[:2]) // 존나신기
}
