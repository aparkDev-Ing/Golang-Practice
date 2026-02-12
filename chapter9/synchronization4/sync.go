package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0

	mutex := new(sync.RWMutex)

	go func() {
		for i := 0; i < 10; i++ {
			mutex.Lock()
			data++
			fmt.Println("Goroutine|writer|", data)
			time.Sleep(200 * time.Millisecond)
			mutex.Unlock()
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			mutex.RLock()
			fmt.Println("Goroutine|reader|", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			mutex.RLock()
			fmt.Println("Goroutine|reader|", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second)
}
