package main

import (
	"fmt"
	"runtime"
)

func inc(v int) int {

	for i := 0; i < 1000; i++ {
		v += i
	}

	return v
}

type Count struct {
	count int
}

func (c *Count) increment() {
	c.count++
}

func (c Count) computeTotal() int {

	return c.count
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	c := Count{
		0,
	}

	done := make(chan bool)

	func() {
		for i := 0; i < 10000; i++ {
			go func() {
				c.increment()
				done <- true
				runtime.Gosched()
			}()
		}
	}()

	for i := 0; i < 10000; i++ {
		<-done
	}

	fmt.Println(c.computeTotal())
}
