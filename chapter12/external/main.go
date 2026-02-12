package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {

	file1 := "sample.xlsx"

	file, err := xlsx.OpenFile(file1)

	errCheck(err)

	for _, sheet := range file.Sheet {

		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				//text := cell.String()
				fmt.Printf("%s\t", cell.String())
			}
			fmt.Println()
		}
	}

}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
