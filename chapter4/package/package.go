package main

import (
	"fmt"
	"os"

	checkUp "github.com/aparkDev-Ing/Golang-Practice/chapter4/lib"
)

func main() {

	var name string

	fmt.Println("Your name: ")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)

	fmt.Println(checkUp.CheckNum(15))

	//fmt.Println(lib2.CheckSum(15))
}
