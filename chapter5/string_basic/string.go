package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var path1 string = "/Users/aaronpark"

	var path2 string = `/Users/aaronpark`

	fmt.Println("path1:", path1)
	fmt.Println("path2:", path2)

	var uniCode rune = 65

	fmt.Printf("%c\n", uniCode)

	//12:28

	var word1 string = "hello"
	var word2 string = "안녕"

	fmt.Println("English byte: ", len(word1))
	fmt.Println("한글바이트: ", len(word2))

	fmt.Println("한글갯구: ", utf8.RuneCountInString(word2))
	fmt.Println("한글갯구: ", len([]rune(word2)))

	//fmt.Println(word1[0])

	fmt.Printf("%c %c %c %c %c\n", word1[0], word1[1], word1[2], word1[3], word1[4])
	fmt.Printf("%c\n", word2[0])

	runeWord2 := []rune(word2)
	fmt.Printf("%c\n", runeWord2[0])

	for _, v := range word1 {
		fmt.Println(string(v))
	}

	runes := []rune(word2)

	for i := 0; i < len(runes); i++ {
		fmt.Println(i, string(runes[i]))
	}

	fmt.Println()

	var a string = "aaron"
	var b string = "애론"

	// for _, v := range a {
	// 	fmt.Printf("%c", v)
	// }
	// fmt.Println()

	// for _, v := range b {
	// 	fmt.Printf("%c", v)
	// }

	for _, v := range a {
		fmt.Print(string(v))
	}
	fmt.Println()

	for _, v := range b {
		fmt.Print(string(v))
	}

	fmt.Println()
}
