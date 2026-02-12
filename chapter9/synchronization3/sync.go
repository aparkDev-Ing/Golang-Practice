package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0

	go func() {
		for i := 0; i < 10; i++ {
			data++
			fmt.Println("Goroutine|writer|", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Goroutine|reader|", data)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Goroutine|reader|", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(5 * time.Second)
}
