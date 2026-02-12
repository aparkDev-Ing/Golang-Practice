package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var v int = 0

func printTest(label string, countDownLatch *sync.WaitGroup) {
	defer countDownLatch.Done()

	startTime := time.Now()
	fmt.Println("Thread:", label, "| Print1 starts")
	//fmt.Println(label, "starts at", start.Format(time.RFC3339Nano))
	time.Sleep(time.Second)
	fmt.Println("Thread:", label, "| Print1 ends")
	endTime := time.Now()
	fmt.Println("Thread:", label, "| Time it took to finish:", endTime.Sub(startTime))
}

func printTest2(label string, ctl *sync.WaitGroup) {
	defer ctl.Done()

	startTime := time.Now()
	fmt.Println("Thread starts:", label)
	for i := 0; i < 1000; i++ {
		fmt.Println("Thread: ", label, "value:", i)
	}
	endTime := time.Now()
	fmt.Println("Thread ends", label, " | Time it took:", endTime.Sub(startTime))
}

func printTestRaceCondition(label string, ctl *sync.WaitGroup) {
	defer ctl.Done()

	startTime := time.Now()
	fmt.Println("Thread starts:", label)
	for ; v < 1001; v++ {
		fmt.Println("Thread: ", label, "value:", v)
	}
	endTime := time.Now()
	fmt.Println("Thread ends", label, "final value of i:", v, " | Time it took:", endTime.Sub(startTime))
}

func exec(label string, ctl *sync.WaitGroup) {
	defer ctl.Done()

	startTime := time.Now()
	randNum := rand.Intn(100)

	for i := 0; i < 100; i++ {
		fmt.Println("Thread name: ", label, "value:", randNum)
	}
	endTime := time.Now()
	fmt.Println("Thread ends", label, " | Time it took:", endTime.Sub(startTime))
}

func closureTest() {

	for i := 0; i < 1000; i++ {
		func(n int) {
			go fmt.Println("Thread name: ", strconv.Itoa(i), "value:", i)
		}(0)
	}
}

func closureTestRaceCondition(a *sync.WaitGroup) int {

	i := 0
	for v := 0; v < 10000000; v++ {
		go func() {
			defer a.Done()
			i++
			//fmt.Println("Thread name: ", strconv.Itoa(v), "value:", v)
		}()
	}

	a.Wait()
	return i
}

func main() {

	// var countDownLatch sync.WaitGroup
	// countDownLatch.Add(2)

	// go printTest("goroutine1", &countDownLatch)

	// printTest("main", &countDownLatch)

	// countDownLatch.Wait()

	// fmt.Println("===============================================")

	// var ctl sync.WaitGroup

	// ctl.Add(2)

	// go printTest2("GoRoutine1", &ctl)
	// printTest2("main", &ctl)

	// ctl.Wait()

	// fmt.Println("===============================================")

	// var ctl sync.WaitGroup

	// ctl.Add(2)

	// go printTest3RaceCondition("GoRoutine1", &ctl)
	// printTest3RaceCondition("main", &ctl)

	// ctl.Wait()

	// fmt.Println("===============================================")

	// //fmt.Println(runtime.NumCPU())
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// //fmt.Println(runtime.GOMAXPROCS(0))

	// var ctl sync.WaitGroup

	// numOfThreads := 100
	// ctl.Add(numOfThreads)

	// for i := 0; i < numOfThreads; i++ {
	// 	go exec(strconv.Itoa(i), &ctl)
	// }

	// //exec("main", &ctl)
	// ctl.Wait()

	// fmt.Println("===============================================")

	// runtime.GOMAXPROCS(1)

	// for i := 0; i < 1; i++ {
	// 	closureTest()
	// }

	// //exec("main", &ctl)
	// time.Sleep(2 * time.Second)

	fmt.Println("===============================================")

	runtime.GOMAXPROCS(1)

	var a sync.WaitGroup

	a.Add(10000000)
	total := closureTestRaceCondition(&a)

	//exec("main", &ctl)

	fmt.Println("total:", total)

}
