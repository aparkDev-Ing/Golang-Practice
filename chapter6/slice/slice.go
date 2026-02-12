package main

import (
	"fmt"
	"sort"
)

func main() {

	intSlice1 := []int{1, 2, 3}

	intSlice2 := make([]int, 0, 3)
	intSlice2 = append(intSlice2, 1, 2, 3)

	intSlice3 := [][]int{
		{1, 2},
		{1, 2},
	}

	intSlice4 := []int{}

	var intSlice5 []int = make([]int, 0, 5)
	intSlice5 = append(intSlice5, 100, 200, 300)

	// var intSlice6 []int
	// intSlice7:= []int{}
	// intSlice8:= make([]int,0)

	//intSlice4 = append(intSlice3, 1, 2, 3)
	// intSlice4[0] = 1
	// intSlice4[1] = 2
	// intSlice4[2] = 3

	//fmt.Println(intSlice1)
	//fmt.Println(intSlice2)
	//fmt.Println(intSlice3)

	fmt.Printf("len: %d cap: %d array: %v\n", len(intSlice1), cap(intSlice1), intSlice1)
	fmt.Printf("len: %d cap: %d array: %v\n", len(intSlice2), cap(intSlice2), intSlice2)
	fmt.Printf("len: %d cap: %d array: %v\n", len(intSlice3), cap(intSlice3), intSlice3)
	fmt.Printf("len: %d cap: %d array: %v\n", len(intSlice4), cap(intSlice4), intSlice4)
	fmt.Printf("len: %d cap: %d array: %v\n", len(intSlice5), cap(intSlice5), intSlice5)

	fmt.Println("------------------------------------------------------------------------")

	intSlice9 := []int{1, 2, 3, 4, 5}
	intSlice10 := intSlice9
	intSlice9[0] = 100

	//밑의 경우 슬라이스헤더의 주소를찍기때문에 다르게나온다 (각개별 다른 헤더를 가진다 복사할때)
	//fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice9), cap(intSlice9), &intSlice9, intSlice9)
	//fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice10), cap(intSlice10), &intSlice10, intSlice10)

	//이렇게 해줘야 backing array address가 나온다
	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice9), cap(intSlice9), &intSlice9[0], intSlice9)
	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice10), cap(intSlice10), &intSlice10[0], intSlice10)

	// for i, v := range intSlice9 {
	// 	fmt.Println("Index:", i, "value:", v)
	// }

	fmt.Println("------------------------------------------------------------------------")

	intSlice11 := [5]int{1, 2, 3, 4, 5}
	intSlice12 := intSlice11
	intSlice11[0] = 100

	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice11), cap(intSlice11), &intSlice11, intSlice11)
	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice12), cap(intSlice12), &intSlice12, intSlice12)

	fmt.Println("------------------------------------------------------------------------")

	//intSlice13 := []int{1, 2, 3, 4, 5}
	intSlice13 := make([]int, 5, 10)
	intSlice13[0] = 1
	intSlice13[1] = 2
	intSlice13[2] = 3
	intSlice13[3] = 4
	intSlice13[4] = 5
	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice13), cap(intSlice13), &intSlice13[0], intSlice13)

	intSlice13 = append(intSlice13, 6, 5, 7)
	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice13), cap(intSlice13), &intSlice13[0], intSlice13)

	fmt.Println("------------------------------------------------------------------------")
	//intSlice13 := []int{1, 2, 3, 4, 5}
	intSlice14 := make([]int, 5, 5)
	intSlice14[0] = 1
	intSlice14[1] = 2
	intSlice14[2] = 3
	intSlice14[3] = 4
	intSlice14[4] = 5
	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice14), cap(intSlice14), &intSlice14[0], intSlice14)

	intSlice14 = append(intSlice14, 6, 5, 7)
	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice14), cap(intSlice14), &intSlice14[0], intSlice14)

	//intSlice13=append(intSlice13,6,5,7)

	fmt.Println("------------------------------------------------------------------------")

	intSlice15 := []int{1, 2, 3, 4, 5}
	intSlice16 := []int{6, 7}
	intSlice17 := []int{8, 9, 10}

	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice15), cap(intSlice15), &intSlice15[0], intSlice15)
	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice16), cap(intSlice16), &intSlice16[0], intSlice16)

	intSlice16 = append(intSlice15, intSlice16...)
	intSlice17 = append(intSlice16, intSlice17[0:2]...)

	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice16), cap(intSlice16), &intSlice16[0], intSlice16)

	fmt.Printf("len: %d cap: %d memory: %p array: %v\n", len(intSlice17), cap(intSlice17), &intSlice17[0], intSlice17)

	fmt.Println("------------------------------------------------------------------------")

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice1[5:])
	fmt.Println(slice1[5:10])
	fmt.Println(slice1[:7])
	fmt.Println(slice1[5:len(slice1)])

	fmt.Println("------------------------------------------------------------------------")

	slice2 := []int{8, 2, 5, 1, 3, 9, 10, 6, 7}
	fmt.Println("Is sorted:", sort.IntsAreSorted(slice2))
	sort.Ints(slice2)
	fmt.Println(slice2)

	slice3 := []string{"b", "d", "a", "e"}
	fmt.Println("Is sorted:", sort.StringsAreSorted(slice3))
	sort.Strings((slice3))
	fmt.Println(slice3)
	fmt.Println("Is sorted:", sort.StringsAreSorted(slice3))

	fmt.Println("------------------------------------------------------------------------")

	slice4 := make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		x := i + 1
		slice4 = append(slice4, x)
	}

	slice5 := make([]int, 5, 5)
	//slice6 := []int{}

	copy(slice5, slice4)
	//copy(slice6, slice4)

	slice5[0] = 7
	slice5[4] = 10

	fmt.Printf("len: %d cap: %d memory: %p slice: %v\n", len(slice4), cap(slice4), &slice4[0], slice4)
	fmt.Printf("len: %d cap: %d memory: %p slice: %v\n", len(slice5), cap(slice5), &slice5[0], slice5)
	//fmt.Println(slice5)
	//fmt.Println(slice6)

	slice6 := []int{1, 2, 3, 4, 5}
	slice7 := slice6

	fmt.Printf("len: %d cap: %d memory: %p slice: %v\n", len(slice6), cap(slice6), &slice6[0], slice6)
	fmt.Printf("len: %d cap: %d memory: %p slice: %v\n", len(slice7), cap(slice7), &slice7[0], slice7)

	fmt.Println("------------------------------------------------------------------------")

	slice8 := [5]int{1, 2, 3, 4, 5}
	slice9 := slice8

	slice9[0] = 0
	fmt.Println(slice8)
	fmt.Println(slice9)

	fmt.Println("------------------------------------------------------------------------")

	slice10 := [5]int{1, 2, 3, 4, 5}
	slice11 := slice10[0:2]

	slice11[0] = 0
	fmt.Println(slice10)
	fmt.Println(slice11)

	fmt.Println("------------------------------------------------------------------------")

	slice12 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice13 := slice12[0:3:5]

	fmt.Println("Slice12: ", slice12)
	fmt.Println("Slice13: ", "cap:", cap(slice13), "len:", len(slice13), "array:", slice13)
}
