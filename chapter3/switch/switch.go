package main

import (
	"fmt"
	//"math/rand"
)

func main() {

	// a := 10

	// fmt.Println(a)

	// switch {

	// case a >= 5:
	// 	fmt.Println("value is greater than or equal to 5: ", a)
	// case a < 5:
	// 	fmt.Println("value is less than 5 ", a)
	// case a%5 == 0:
	// 	fmt.Println("value is even ")
	// case a%5 > 0:
	// 	fmt.Println("value is odd")
	// }

	// switch a {

	// case 5:
	// 	fmt.Println("value is: ", a)
	// }

	//using local variable inside switch
	// switch a := "hello"; a {

	// case "hello":
	// 	fmt.Println("hello")
	// case "bye":
	// 	fmt.Println("bye")
	// case "hi":
	// 	fmt.Println("hi")
	// case "sayonara":
	// 	fmt.Println("sayonara")

	// default:
	// 	fmt.Println("default value")
	// }

	//fmt.Println(a)

	// switch a := "hello"; a + " go" {

	// case "hello java":
	// 	fmt.Println("java")
	// case "hello go":
	// 	fmt.Println("go")
	// case "hello python":
	// 	fmt.Println("python")
	// case "hello ruby":
	// 	fmt.Println("ruby")

	// default:
	// 	fmt.Println("default value")
	// }

	// switch a, b := "hello", "hi"; {

	// case a == b:
	// 	fmt.Println("a is equal to b")

	// default:
	// 	fmt.Println("default value")
	// }

	switch a := "hello"; {

	case a == "hello java":
		fmt.Println("java")
	case a == "hello":
		fmt.Println("hello")
	default:
		fmt.Println("default value")
	}

	//fmt.Println(rand.Intn(100))

	// switch i := rand.Intn(100); {
	// case i >= 0 && i < 51:
	// 	fmt.Println("I is between 0 and 50")

	// case i >= 51 && i < 76:
	// 	fmt.Println("I is between 51 and 75")

	// case i >= 76 && i < 100:
	// 	fmt.Println("I is between 76 and 100")

	// }

	// var (
	// 	a int = 15
	// 	b int = 30
	// 	c int = 45
	// )

	// switch {
	// case a > 10, b > 10, c > 10:
	// 	fmt.Println("a|b|c is greater than 10. a b c in order", a, b, c)

	// }

	switch e := "go"; e {
	case "java":
		fmt.Println("java")
		fallthrough
	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
		fallthrough
	case "c":
		fmt.Println("c")
		fallthrough
	case "ruby":
		fmt.Println("ruby")
	}

}
