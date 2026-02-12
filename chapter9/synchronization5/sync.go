package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan int, 5)
	mutex := new(sync.Mutex)
	condition := sync.NewCond(mutex)

	for i := 0; i < 5; i++ {
		go func(v int) {
			mutex.Lock()
			c <- v
			fmt.Println("Goroutine Thread | waits...")
			condition.Wait()
			fmt.Println("Goroutine Thread | wakes up...")
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Main received:", <-c)
	}

	//one by one
	// for i := 0; i < 5; i++ {
	// 	mutex.Lock()
	// 	fmt.Println("Main Thread | signals...")
	// 	condition.Signal()
	// 	mutex.Unlock()
	// }

	//at once
	// mutex.Lock()
	// condition.Broadcast()
	// mutex.Unlock()

	time.Sleep(3 * time.Second)
}
