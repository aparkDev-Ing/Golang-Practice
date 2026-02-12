package main

import (
	"fmt"
	"os"
)

func errorCheck(e error) {

	if e != nil {
		fmt.Println("Error Occurred:", e)
		panic(e)
	}

}

func main() {

	file, err := os.Create("files/test.txt")
	errorCheck(err)

	defer file.Close()

	s1 := []byte{1, 2, 3, 4, 5, 7, 8}

	n1, err := file.Write(s1)

	errorCheck(err)

	fmt.Println("File size: ", n1)

	fmt.Println("===============================")

	n2, err := file.WriteString("\nhello aaron")

	file.Sync()

	errorCheck(err)

	fmt.Println("File size: ", n2)

	fmt.Println("===============================")

	n3, err := file.WriteAt([]byte("\nhello aaron"), 70)

	errorCheck(err)

	fmt.Println("File size: ", n3)
}
