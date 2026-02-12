package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan string)

	//수신 애플리케이션
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Integer", <-c1)
		}
	}()

	//발신 애플리케이션
	go func() {
		for {
			value := "Hello"
			time.Sleep(1 * time.Second)
			//fmt.Println("Producer Thread | String", value)
			c2 <- value
		}
	}()

	//수발신 모두
	go func() {

		for {
			select {
			case c1 <- rand.Intn(100):
			case str := <-c2:
				fmt.Println("Consumer Thread | String", str)
				// default:
				// 	fmt.Println("Default")
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
