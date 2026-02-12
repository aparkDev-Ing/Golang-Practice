package main

import (
	"fmt"
)

func main() {

	s := "hello"
	v := "안녕"

	//returns byte size
	fmt.Println("hello size:", len(s))
	fmt.Println("안녕크기:", len(v))

	fmt.Println("===========================")

	//returns first index of byte
	fmt.Println("first digit:", s[0])
	fmt.Println("첫번째글짜:", v[0])

	fmt.Println("===========================")
	//ascii matches unicode codepoint
	fmt.Printf("first digit:%c\n", s[0])
	//utf-8 does not match with unicode codepoint
	fmt.Printf("첫번째글짜:%c\n", v[0])

	fmt.Println("===========================")

	//unicode codepoint
	fmt.Printf("first digit:%c\n", []rune(s)[0])
	//unicode codepoint
	fmt.Printf("첫번째글짜:%c\n", []rune(v)[0])
}
