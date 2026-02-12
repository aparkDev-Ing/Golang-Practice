package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			value := rand.Intn(100)
			time.Sleep(1 * time.Second)
			//fmt.Println("Producer Thread | Integer", value)
			c1 <- value
		}
	}()

	go func() {
		for {
			value := "Hello"
			time.Sleep(1 * time.Second)
			//fmt.Println("Producer Thread | String", value)
			c2 <- value
		}
	}()

	//수신처리
	go func() {

		for {
			select {
			case num := <-c1:
				fmt.Println("Consumer Thread | Integer", num)
			case str := <-c2:
				fmt.Println("Consumer Thread | String", str)
				// default:
				// 	fmt.Println("Default")
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
