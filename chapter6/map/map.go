package main

import (
	"fmt"
)

func main() {

	// basic way to initilize
	var map1 map[string]int = make(map[string]int)

	var map2 = make(map[string]int)

	map3 := make(map[string]int)

	fmt.Println("ex :", map1)
	fmt.Println("ex :", map2)
	fmt.Println("ex :", map3)

	map4 := map[string]int{}

	map4["apple"] = 25
	map4["banana"] = 40
	map4["organge"] = 60

	map5 := map[string]int{
		"apple":  25,
		"banana": 25,
		"oragne": 25,
	}

	map6 := map[string]int{"apple": 25, "banana": 40, "oragne": 60}

	map6["grape"] = 70

	map7 := make(map[string]int, 5)
	map7["apple"] = 15
	map7["grape"] = 15
	map7["oragne"] = 15
	map7["banana"] = 15
	map7["kiwi"] = 15
	map7["strawberry"] = 15

	fmt.Println(map4)
	fmt.Println(map5)
	fmt.Println(map6)
	fmt.Println(map7)

	fmt.Println("-----------------------------------------------------")

	map8 := map[string]string{
		"Mexico":  "Taco",
		"Korea":   "Bibimbap",
		"America": "Burger",
	}

	//fmt.Println(map8["Mexico"])
	for k, v := range map8 {
		fmt.Println("Key: ", k, "Value:", v)
	}

	for _, v := range map8 {
		fmt.Println("value:", v)
	}

	for k := range map8 {
		fmt.Println("key:", k)
	}

	fmt.Println("-----------------------------------------------------")

	map9 := make(map[string]int, 10)

	map9["Aaron"] = 70
	map9["Bj"] = 65
	map9["Sejun"] = 80
	map9["random"] = 100
	map9["Zero"] = 0

	delete(map9, "random")

	map9["jaeyoung"] = 90
	fmt.Println(map9)

	fmt.Println("-----------------------------------------------------")

	value1 := map9["Aaron"]
	value2 := map9["Bj"]
	value3, ok1 := map9["None"]
	value4, ok2 := map9["Zero"]

	fmt.Println(value1, value2, value3, ok1, value4, ok2)

	if ok2 {
		fmt.Println("Zero Exists as entry in map")
	}

	if value, ok := map9["BJ"]; ok {
		fmt.Println("Bj is in map. Value:", value)
	} else {
		fmt.Println("Does not exist")
	}

	if _, ok3 := map9["David"]; !ok3 {
		fmt.Println("David does not exist", ok3)
	}

}
