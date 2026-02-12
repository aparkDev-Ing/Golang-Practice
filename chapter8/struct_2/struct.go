package main

import "fmt"

type Athele struct {
	name      string "이름"
	experiece int
}

type UFCFighter struct {
	Athele
	//name  string
	style string
}

func (f UFCFighter) fight() {
	fmt.Println("Ufc fight starts")
}

func (f Athele) fight() {
	fmt.Println("Athlete fight starts")
}

func newCorporate(name string, kind string, revenue float64, expense float64, totalIncome float64) *Corporate {
	return &Corporate{
		name,
		kind,
		revenue,
		expense,
		totalIncome,
	}
}

type Corporate struct {
	name        string
	kind        string
	revenue     float64
	expense     float64
	totalIncome float64
}

const (
	tier3 float64 = 100000
	tier2 float64 = 200000
	tier1 float64 = 300000
)

// non receiver method - must require pointer type
// func computeIncome(c *Corporate) float64 {

// 	fmt.Printf("c's address:%p\n", &c)
// 	total := c.revenue - c.expense
// 	c.totalIncome = total
// 	return total
// }

func (c *Corporate) computeIncome() float64 {

	fmt.Printf("c's address:%p\n", &c)
	total := c.revenue - c.expense
	c.totalIncome = total
	return total
}

// func (c Corporate) computeIncome() float64 {

// 	fmt.Printf("c's address:%p\n", &c)
// 	total := c.revenue - c.expense
// 	c.totalIncome = total
// 	return total
// }

func main() {

	abc := Corporate{
		"abc",
		"tech",
		100000,
		5000,
		0,
	}

	bcd := &Corporate{
		"bcd",
		"hospitliaty",
		200000,
		100000,
		0,
	}

	apple := newCorporate("apple", "tech", 100000000, 200000, 0)
	fmt.Println("abc:", abc)
	fmt.Println("bcd", bcd)

	fmt.Println("===================================================================")

	fmt.Println(apple)
	fmt.Printf("apple's address:%p\n", &apple)
	//fmt.Println("apple's revenue:", int(apple.revenue), "apple's revenue:", int(computeIncome(apple)))
	fmt.Println("apple's revenue:", int(apple.revenue), "apple's revenue:", int(apple.computeIncome()))
	fmt.Println("apple's total income:", int((*apple).totalIncome))

	fmt.Println("===================================================================")

	jones := UFCFighter{

		Athele{
			"jon jones",
			5,
		},
		"Freestyle",
	}
	fmt.Println(jones)

	jones.fight()

	fmt.Println("===================================================================")

	jones.Athele.fight()
}
