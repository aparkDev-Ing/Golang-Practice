package main

import (
	"fmt"
	"math/rand"
	"time"
)

func exec1(v chan int) {
	startTime := time.Now()
	time.Sleep(1 * time.Second)
	endTime := time.Now()
	fmt.Println("Time it took to finish exec1:", endTime.Sub(startTime))
	v <- 1
}

func exec2(v chan int) {
	startTime := time.Now()
	time.Sleep(1 * time.Second)
	endTime := time.Now()
	fmt.Println("Time it took to finish exec2:", endTime.Sub(startTime))
	v <- 2
}

func rangeSum(v chan int, k int) {

	sum := 0

	for i := 0; i < k; i++ {
		sum += i
	}

	v <- sum
}

func print(ch chan int) {
	for v := range ch {
		fmt.Println("value:", v)
	}
}

func main() {
	// v := make(chan int)

	// go exec1(v)
	// go exec2(v)

	// <-v
	// <-v

	// v := make(chan int)

	// go rangeSum(v, 500)
	// go rangeSum(v, 1000)
	// go rangeSum(v, 1500)

	// result1 := <-v
	// result2 := <-v
	// result3 := <-v

	// fmt.Println("result1:", result1)
	// fmt.Println("result2:", result2)
	// fmt.Println("result3:", result3)

	// c := make(chan bool)
	// cnt := 0
	// go func() {
	// 	for i := cnt; i < 6; i++ {
	// 		c <- true
	// 		fmt.Println("go thread: ", i)
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	// for i := cnt; i < 6; i++ {
	// 	<-c
	// 	fmt.Println("main thread: ", i)
	// }

	// runtime.GOMAXPROCS(1)

	// c := make(chan bool, 2)
	// cnt := 0

	// go func() {
	// 	for i := cnt; i < 5; i++ {
	// 		c <- true
	// 		fmt.Println("goroutine:", i)
	// 	}
	// }()

	// for i := cnt; i < 5; i++ {
	// 	<-c
	// 	fmt.Println("main:", i)
	// }

	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Goroutine sequence:", i+1)
			ch <- rand.Intn(100)
		}
		//close(ch)
	}()

	fmt.Println("Main starts to print display")
	//Program errors out with deadlock because once sender finishes send, there is other process who keeps sending and main thread keeps blockign waiting on receive
	print(ch)

	// ch := make(chan int)

	// go exec2(ch)

	// value1, ok1 := <-ch
	// if ok1 {
	// 	fmt.Println("value1:", value1)
	// }

	// close(ch)

	// value2, ok2 := <-ch
	// if !ok2 {
	// 	fmt.Println("value2:", value2)
	// }
}
