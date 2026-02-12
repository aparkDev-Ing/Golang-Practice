package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func errorCheck(e error) {

	if e != nil {
		fmt.Println("Error Occurred:", e)
		panic(e)
	}

}

func main() {

	file, err := os.Open("files/sample.csv")
	errorCheck(err)
	defer file.Close()

	//fileStat, err := file.Stat()
	errorCheck(err)

	rr := csv.NewReader(bufio.NewReader(file))

	row1, err1 := rr.Read()
	row2, err2 := rr.Read()
	errorCheck(err1)
	errorCheck(err2)

	fmt.Println("Row1 data:", row1)
	fmt.Println("Row1 data Sliced:", row1[1:5])
	fmt.Println("Row2 data:", row2)
	fmt.Println("Row2 data Sliced:", row2[1:5])
	// fmt.Println("===========================================")

	rows, err := rr.ReadAll()
	errorCheck(err)

	for i, v := range rows {
		fmt.Println("row", strconv.Itoa(i+3), "value:", v)
	}

	// for i, v := range rows {
	// 	for c, v := range v {
	// 		fmt.Println("Row:", i+3, "Column:", c+1, "value:", v)
	// 	}
	// }

	// for i := 0; i < 30; i++ {
	// 	for c, v := range rows[i] {
	// 		fmt.Println("Row:", i+3, "Column:", c+1, "value:", v)
	// 	}
	// }

	fmt.Println("===========================================")

	// rows, err := rr.ReadAll()
	// errorCheck(err)

	// row3Region := rows[0][0]
	// fmt.Println("Row3 Region", row3Region)

	// fmt.Println("파일정보:", fileStat)
	// fmt.Println("파일이름:", fileStat.Name())
	// fmt.Println("파일사이즈:", fileStat.Size())

}
