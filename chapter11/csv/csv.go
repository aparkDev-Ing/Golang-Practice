package main

import (
	_ "bufio"
	"encoding/csv"
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

	f1, err := os.Create("files/test.csv")
	errorCheck(err)
	defer f1.Close()

	wr := csv.NewWriter(f1)

	//wr1:=csv.NewWriter(bufio.NewWriter(f1))

	wr.Write([]string{"aaron", "1.8"})
	wr.Flush()

	fi, err := f1.Stat()

	errorCheck(err)
	fmt.Println(fi.Size())
	fmt.Println(fi.Name())
	fmt.Println(fi.Mode())
}
