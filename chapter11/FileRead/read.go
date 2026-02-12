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

	f1, err := os.Open("files/test.txt")

	errorCheck(err)

	fileInfo, err := f1.Stat()
	errorCheck(err)
	sli1 := make([]byte, fileInfo.Size())
	ct1, err := f1.Read(sli1)

	fmt.Println("파일정보:", fileInfo)
	fmt.Println("파일이름:", fileInfo.Name())
	fmt.Println("파일사이즈:", fileInfo.Size())
	fmt.Println("byte size:", ct1)
	fmt.Println("Data raw bytes[Will be printed in ASCII]:", sli1)
	fmt.Println("Data:", string(sli1))

	fmt.Println("======================================================")

	o2, err := f1.Seek(20, 0)
	errorCheck(err)

	sli2 := make([]byte, 30)
	ct2, err := f1.Read(sli2)
	errorCheck(err)

	fmt.Println("파일정보:", fileInfo)
	fmt.Println("파일이름:", fileInfo.Name())
	fmt.Println("파일사이즈:", fileInfo.Size())
	fmt.Println("byte size:", ct2)
	fmt.Println("Data raw bytes[Will be printed in ASCII]:", sli2)
	fmt.Println("Data:", string(sli2))

	fmt.Println("o1:", o2)
	fmt.Println("======================================================")

	o3, err := f1.Seek(0, 0)
	errorCheck(err)

	sli3 := make([]byte, 50)
	ct3, err := f1.ReadAt(sli3, 8)
	errorCheck(err)

	fmt.Println("파일정보:", fileInfo)
	fmt.Println("파일이름:", fileInfo.Name())
	fmt.Println("파일사이즈:", fileInfo.Size())
	fmt.Println("byte size:", ct3)
	fmt.Println("Data raw bytes[Will be printed in ASCII]:", sli3)
	fmt.Println("Data:", string(sli3))

	fmt.Println("o3:", o3)
	fmt.Println("======================================================")

	defer f1.Close()
}
