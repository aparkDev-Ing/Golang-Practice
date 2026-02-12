package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("unknownfile")

	if err != nil {
		//fmt.Println(err)
		fmt.Println("file name:", f)
		log.Fatal(err.Error())
		//log.Fatal(err)
	}

	fmt.Println("file name:", f.Name())

}
