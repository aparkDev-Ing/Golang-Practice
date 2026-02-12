package main

import (
	"fmt"
	"reflect"
)

func main() {

	print(1)
	print("Hello")
	print([]int{1, 2, 3, 4, 5})
	fmt.Println("----------------------------------------")
	printVal(1)
	printVal("Hello")
	printVal([]int{1, 2, 3, 4, 5})
	fmt.Println("----------------------------------------")

	var a interface{} = 15

	b := a
	c := a.(int) //using assertion to extract int value
	//d := a.(float64)

	fmt.Println("value: ", a, "Type:", reflect.TypeOf(a))
	fmt.Println("value: ", b, "Type:", reflect.TypeOf(b))
	fmt.Println("value: ", c, "Type:", reflect.TypeOf(c))

	fmt.Println("----------------------------------------")

	if v, ok := a.(int); ok {
		fmt.Println("type:", reflect.TypeOf(v), "value", v)
	}

	fmt.Println("----------------------------------------")

}

func print(behavior interface{}) {
	fmt.Printf("Type: (%T) ", behavior)
	//fmt.Println(reflect.TypeOf(behavior))
	fmt.Println("| Value: ", behavior)
}

func printVal[T any](v T) {
	fmt.Printf("Type: (%T) ", v)
	fmt.Println("| Value: ", v)
}

func checkTypeInter(behavior interface{}) {

	switch behavior.(type) {
	case bool:
		fmt.Println("this is bool:", behavior)
	case uint8, int, int8, int16, int32, int64:
		fmt.Println("this is int:", behavior)
	case string:
		fmt.Println("this is string:", behavior)

	}

}
