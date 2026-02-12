package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var mutex sync.Mutex

func compute(total *int64, count int, countDownLatch *sync.WaitGroup) {

	for i := 0; i < count; i++ {
		//countDownLatch.Add(1)
		go func() {
			defer countDownLatch.Done()
			for i := 0; i < 100; i++ {
				atomic.AddInt64(total, 1)
			}
		}()
	}
}

func main() {

	//var countDownLatch sync.WaitGroup

	countDownLatch := new(sync.WaitGroup)

	var total int64 = 0
	const count = 10
	countDownLatch.Add(count)
	compute(&total, count, countDownLatch)
	countDownLatch.Wait()

	//fmt.Println(total)
	fmt.Println(atomic.LoadInt64(&total))
}
