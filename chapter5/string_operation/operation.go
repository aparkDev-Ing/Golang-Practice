package main

import (
	"fmt"
	"strings"
)

func main() {

	var str1 string = "hello"
	var str2 string = "world"
	var str3 string = "world"

	fmt.Println(str1[0:2])
	fmt.Println(str2[3:])
	fmt.Println(str2[3:])
	fmt.Println(str2[:4])
	fmt.Println(str2[1:3])

	fmt.Println(str1 == str2)
	fmt.Println(str2 == str3)

	var str4 string = "aaron"
	var str5 string = "kevin"

	fmt.Println(str4 > str5)

	// fmt.Println(str4 > str5) 가 왜 false냐?

	// Go의 문자열 비교(<, >, <=, >=)는 사전식(lexicographical) 비교이고, 기준은 바이트 값(UTF-8 바이트) 순서입니다.

	// str4 := "aaron"
	// str5 := "kevin"
	// fmt.Println(str4 > str5) // false

	// 비교는 앞에서부터:

	// 'a' vs 'k'

	// ASCII에서 'a'(97) < 'k'(107) 이므로 "aaron" < "kevin"

	// 따라서 "aaron" > "kevin" 은 false

	// ✅ 결론: 네가 말한 bash처럼 배쉬도 마찬가지 “앞글자부터 알파벳 순서” 느낌이 맞고, Go는 그걸 “바이트(UTF-8) 기준 사전식 비교”로 한다고 보면 돼요.

	str6 := "네가 말한 bash처럼 배쉬도 마찬가지 앞글자부터 알파벳 순서 느낌이 맞고, Go는 그걸 바이트(UTF-8) 기준 사전식 비교로 한다고 보면 돼요." +
		"네가 말한 bash처럼 배쉬도 마찬가지 “앞글자부터 알파벳 순서” 느낌이 맞고, Go는 그걸 “바이트(UTF-8) 기준 사전식 비교”로 한다고 보면 돼요." +
		"네가 말한 bash처럼 배쉬도 마찬가지 “앞글자부터 알파벳 순서” 느낌이 맞고, Go는 그걸 “바이트(UTF-8) 기준 사전식 비교”로 한다고 보면 돼요."

	str7 := "네가 말한 bash처럼 배쉬도 마찬가지 “앞글자부터 알파벳 순서” 느낌이 맞고, Go는 그걸 “바이트(UTF-8) 기준 사전식 비교”로 한다고 보면 돼요."

	fmt.Println(str6 + str7)

	strSet := []string{}
	strSet = append(strSet, str6)
	strSet = append(strSet, str7)

	str8 := strings.Join(strSet, "")

	fmt.Println((str6 + str7) == str8)

	// intArr := []rune{1, 2, 3}

	// for _, v := range intArr {
	// 	fmt.Println(v)
	// }

	// intArr := []int{1, 2, 3}

	// fmt.Println(intArr)
}
