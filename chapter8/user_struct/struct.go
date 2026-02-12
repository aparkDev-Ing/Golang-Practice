package main

import (
	"fmt"
)

type Car struct {
	name  string
	color string
	price int
	year  int
}

type Customer struct {
	name       string
	membership string
	points     memberShipPoint
}

type memberShipPoint int

func initialCustomer(c Customer, point memberShipPoint) {
	c.points += point
}

func (c Customer) updateCustomerInfo1(point memberShipPoint) {
	c.points += point
}

func (c *Customer) updateCustomerInfo2(point memberShipPoint) {
	(*c).points += point
}

type cnt int

type totalCost func(total int, cnt cnt) int

func computeGrandTotal(total int, cnt cnt, f totalCost) int {

	return f(total, cnt)
}

func compute(i int, cnt cnt) int {
	return i + 1
}

func Price(c Car) int {
	return c.price
}
func (c Car) Price() int {
	return c.price
}

func defaultType(i int) {
	fmt.Println("Default: ", i)
}

func customType(i cnt) {
	fmt.Println("custom: ", i)
}

func main() {

	bmw := Car{
		name:  "bmwx1",
		color: "white",
		price: 10000,
		year:  2025,
	}

	fmt.Println(bmw)
	fmt.Printf("Car: %v Inventory Address: %p\n", bmw, &bmw)

	fmt.Println(Price(bmw))
	fmt.Println(bmw.Price())

	fmt.Println("---------------------------------------------------------------------------")

	var total int = 100
	var a cnt = 15

	//defaultType(a)
	defaultType(int(a))

	customType(a)

	fmt.Println("---------------------------------------------------------------------------")

	var b totalCost = func(i int, cnt cnt) int {

		return i + int(cnt)
	}

	// var c totalCost= compute

	fmt.Println("Total: ", computeGrandTotal(total, a, b))

	fmt.Println("---------------------------------------------------------------------------")

	customer1 := Customer{
		name:       "aaron",
		membership: "VIP",
		points:     0,
	}

	initialCustomer(customer1, 100)

	customer1.updateCustomerInfo1(100)

	fmt.Println(customer1)

	//value type으로 reference method 를 콜하는모습
	customer1.updateCustomerInfo2(100)

	fmt.Println(customer1)

}
