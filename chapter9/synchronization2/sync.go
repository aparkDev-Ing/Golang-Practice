package main

import (
	"fmt"
	"runtime"
	"sync"
)

func inc(v int) int {

	for i := 0; i < 1000; i++ {
		v += i
	}

	return v
}

type Count struct {
	count int
	mutex sync.Mutex
}

func (c *Count) increment() {
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
}

// 실무관점에선 포인터가 맞다 불변한 데이터를위해서 (스트럭트 값을보호) 근데 사실그냥값으로해도 여기선문제없다
func (c Count) computeTotal() int {

	return c.count
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	c := Count{
		count: 0,
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
	//fmt.Println(c.count)
}
