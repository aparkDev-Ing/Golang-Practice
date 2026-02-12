package main

import (
	"fmt"
	"math/rand/v2"
)

type Behavior interface {
	pay(i int)
}

type PaymentGateway interface {
	payThroughGateway(c Customer)
}

type PayPalGateway struct {
	name string
}

type klaranaGateway struct {
	name string
}

type Customer struct {
	name       string
	customerId int
	age        int
	//paymentType PaymentType
	gateway PaymentGateway
}

type PaymentType struct {
	name string
}

func (c Customer) pay(i int) {
	fmt.Println(c.name, "|Customer pays for an amount of:", i)
	fmt.Println(c.name, "Customer's payment method:", c.gateway, "calling this gateway...")
	c.gateway.payThroughGateway(c)
	fmt.Println("Payment successful!")
}

func (g PayPalGateway) payThroughGateway(c Customer) {
	fmt.Println(c.name, "|", g, "| invoking customer payment gateway ...")
	fmt.Println("Payment gateway call sucessful!")
}

func (g klaranaGateway) payThroughGateway(c Customer) {
	fmt.Println(c.name, "|", g, "| invoking customer payment gateway ...")
	fmt.Println(g.name, "requires more authentication....")
	fmt.Println("Authentication successful...")
	fmt.Println("Payment gateway call sucessful!")
}

func actBehaviors(behaviors []Behavior) {
	for _, v := range behaviors {
		randNum := rand.IntN(100)
		v.pay(randNum)
	}
}

func act(customer interface{ pay(i int) }) {
	randNum := rand.IntN(100)
	customer.pay(randNum)
}

// 데이터타입을 요구하지않는 빈 annomoymous interface를 이용하는 예제
func printVal(s interface{}) {
	fmt.Println(s)
}

func main() {

	// var paypal behavior

	// aaron := Customer{
	// 	"Aaron",1,28,
	// }

	// paypal=aaron

	// fmt.Println(paypal)

	// var aaron behavior = Customer{
	// 	"Aaron", 1, 28,
	// }
	// aaron := behavior(Customer{
	// 	"Aaron", 1, 28,
	// })

	// var david behavior = Customer{
	// 	"david", 1, 28,
	// }

	// aaron.pay()
	// david.pay()

	// var pg1 PaymentGateway = PayPalGateway{
	// 	"paypal",
	// }
	// pg1.payThroughGateway(Customer{})

	fmt.Println("--------------------------------------------------------------------------------------")

	var aaron Behavior = Customer{
		"aaron",
		12345,
		28,
		PayPalGateway{
			"paypal",
		},
	}
	david := Customer{
		"david",
		23456,
		32,
		klaranaGateway{
			"klarana",
		},
	}

	inters := []Behavior{aaron, david}
	//customers := []Customer{aaron, david}

	// for _, v := range inters {
	// 	v.pay(15)
	// }

	actBehaviors(inters)

	//Behavior(aaron).pay(15)
	//Behavior(david).pay(50)

	fmt.Println("--------------------------------------------------------------------------------------")

	act(inters[0])

	fmt.Println("--------------------------------------------------------------------------------------")

	printVal([5]int{12, 3, 4, 5})
	printVal([]int{12, 3, 4, 5})
	printVal(map[string]int{
		"hello": 1,
	})
}
