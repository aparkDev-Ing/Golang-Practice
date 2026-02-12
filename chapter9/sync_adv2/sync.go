package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	var countDownLatch sync.WaitGroup

	fmt.Println("Main thread assigns work...")
	count := 5
	countDownLatch.Add(count)
	calculate(count, &countDownLatch)
	countDownLatch.Wait()
	fmt.Println("Main thread finishes assign work...")

	//time.Sleep(1 * time.Second)
}

func calculate(count int, countDownLatch *sync.WaitGroup) {

	for i := 0; i < count; i++ {
		//이렇게해도된다
		//countDownLatch.Add(1)
		go func() {
			defer countDownLatch.Done()
			count := 0
			numb := rand.Intn(100)
			for i := 0; i < numb; i++ {
				count += i
			}
			fmt.Println("Goroutine | ", i+1, "value:", count)

		}()

	}
}
