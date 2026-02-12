package main

import (
	"fmt"
	"os"
)

func openFile(fileName string) {

	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error Occurred while opening file:", s)
		}
	}()

	f, err := os.Open(fileName)

	if err != nil {
		//log.Panic(err)
		panic(err)
	}

	fmt.Println("filename:", f)
	f.Close()
}

func main() {

	openFile("unknwon")
}
