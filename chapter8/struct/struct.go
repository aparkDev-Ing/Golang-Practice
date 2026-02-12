package main

import (
	"fmt"
	"reflect"
)

type Food struct {
	name    string
	price   int
	cuisine string
}

type Car struct {
	name   string "브랜드"
	year   int    "년도"
	engine engine "엔진"
	logo   struct {
		name string
	}
}

type engine struct {
	name string "엔진이름"
}

// receiver
func (a Food) compute(taxRate float64) float64 {

	return float64(a.price) * taxRate
}

// func (a *Food) compute1(taxRate float64) float64 {

// 	return float64(a.price) * taxRate
// }

func main() {

	burger := Food{
		name:    "Cheese Burger",
		price:   10,
		cuisine: "American",
	}

	//does not work if receiver mothod above requires pointer
	//Food{"Cheeze Pizaa", 20, "Italian"}.compute1(10.0)

	pizza := Food{"Cheeze Pizaa", 20, "Italian"}

	//pizza.compute1(10.0)

	fmt.Println("---------------------------------------------------------------------------")

	//var pasta *Food= new(Food)
	//pasta := new(Food)
	var pasta *Food = new(Food)
	pasta.name = "Alio olio"
	pasta.price = 15
	pasta.cuisine = "Italian"

	taco := &Food{
		name:    "Tripa",
		price:   10,
		cuisine: "Mexican",
	}

	fmt.Println("Pizza: ", pizza)
	fmt.Println("burger: ", burger)
	fmt.Println("taco: ", taco)
	fmt.Println("Pasta: ", pasta)

	fmt.Printf("Taco: %#v\n", taco)
	fmt.Printf("pizza: %#v\n", pizza)

	fmt.Println("---------------------------------------------------------------------------")

	//익명 struct
	sushi := struct {
		name  string
		price int
	}{name: "sushi", price: 15}

	tacoyaki := struct {
		name  string
		price int
	}{"tacoyaki", 5}

	fmt.Println(sushi, tacoyaki)

	foodList := []struct {
		name  string
		price int
	}{{"curry", 20}, {"fried rice", 20}}

	for _, v := range foodList {
		fmt.Println(v)
	}

	fmt.Println("---------------------------------------------------------------------------")

	// car1 := Car{
	// 	"benz", 2015,
	// }

	tag := reflect.TypeOf(Car{})
	for i := 0; i < tag.NumField(); i++ {
		fmt.Println(tag.Field(i))
	}

	fmt.Println("---------------------------------------------------------------------------")

	car1 := Car{
		"Porche",
		250000,
		engine{
			"Yasuo",
		},
		struct{ name string }{"Macan"},
	}
	car2 := Car{
		"Porche",
		250000,
		engine{
			"yasuo",
		},
		struct{ name string }{"Taycan"},
	}
	fmt.Println("car1:", car1)
	fmt.Println("first car's logo: ", car1.logo.name)
	fmt.Println("car2:", car2)
	fmt.Println("second car's logo: ", car2.logo.name)
}
