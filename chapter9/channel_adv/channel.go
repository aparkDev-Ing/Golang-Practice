package main

import (
	"fmt"
	"time"
)

func sendOnly(label string, c chan<- int, count int) {

	for i := 0; i < count; i++ {
		fmt.Println("Producer: ", label, " sends data:", i)
		c <- i
	}

}

func receiveOnly(label string, c <-chan int) {

	//missing close not causing deadlock because child routines are blocked waiting receive but main terminates in few seconds
	for i := range c {
		fmt.Println("Consumer: ", label, " receives data:", i)
	}
}

func produce(count int) <-chan int {
	value := 0
	c := make(chan int)

	go func() {
		fmt.Println("Gothread | Producer | starts calculate numers")
		for i := 0; i < count; i++ {
			value += i
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Gothread | Producer | ends calculate numers")
		c <- value
		//c <- rand.Intn(100)
		//c <- rand.Intn(100)

		close(c)
	}()

	return c
}

func compute(c <-chan int) <-chan int {

	computeChannel := make(chan int)
	sum := 0

	go func() {
		fmt.Println("Gothread | Computer | starts compute total")
		for i := range c {
			sum += i
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Gothread | Computer | ends compute total")
		computeChannel <- sum
	}()

	return computeChannel
}

func main() {
	// c := make(chan int)
	// go sendOnly("Gopro", c, 10)
	// go receiveOnly("Gorec", c)
	// time.Sleep(2 * time.Second)

	fmt.Println("========================================")

	//채널을 안에서만들고 수신채널을 리턴받아서 메인에서 수신받고 출력
	// receiver := produce(100)
	// fmt.Println("Main thread waiting...")
	// final := <-receiver
	// fmt.Println("Main thread finished waiting. Value:", final)

	fmt.Println("========================================")

	//채널을 안에서만들고 수신채널 리턴받고 그걸이용해전체값계산후 또 수신채널 리턴받고 수신받고 출력
	//메인은 두개의채널을만들고 compute채널을 기다린다
	//첫번째스레드는 3가지 번호를 전송하고 끝낸다 (close필수)
	//두번째스레드는 첫번째채널에서 번호를받고 토탈을 compute하고 이걸다시 두번째채널에 전송하고끝

	// valuesReceiver := produce(100)
	// totalReceiver := compute(valuesReceiver)
	// fmt.Println("Main thread waiting for computation")
	// total := <-totalReceiver
	// fmt.Println("Total compute completed: ", total)
	//fmt.Println(<-totalReceiver)

}
