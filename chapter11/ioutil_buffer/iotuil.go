package main

import (
	"fmt"
	"io"
	"os"

	//"ioutil"
	"bufio"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// s := "hello world!"
	// err := os.WriteFile("files/hello world.txt", []byte(s), 0644)
	// errorCheck(err)

	file, err := os.OpenFile("files/hello world.txt", os.O_CREATE|os.O_RDWR, 0644)
	errorCheck(err)

	writeBuffer := bufio.NewWriter(file)

	s := "hello world"
	writeBuffer.WriteString("Hello World")
	writeBuffer.Write([]byte(s))

	//fmt.Println("버퍼:", writeBuffer)
	fmt.Printf("Used Size out of buffer (%d bytes)\n", writeBuffer.Buffered())
	fmt.Printf("Left Buffer Size (%d bytes)\n", writeBuffer.Available())
	fmt.Printf("Total Buffer Size (%d bytes)\n", writeBuffer.Size())

	writeBuffer.Flush()

	fmt.Println("========================================================")

	readBuffer := bufio.NewReader(file)
	fileData, err := file.Stat()
	errorCheck(err)

	data := make([]byte, fileData.Size())
	file.Seek(0, io.SeekStart)

	readData, err := readBuffer.Read(data)
	errorCheck(err)

	//readData, err := readBuffer.Read(data)

	fmt.Println("Data:", string(data))
	fmt.Println("Data Size:", fileData.Size())
	fmt.Println("Data Size:", readData)

	defer file.Close()
}
