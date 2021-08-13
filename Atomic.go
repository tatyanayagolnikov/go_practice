package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/// RACE CONDITION ///


func main() {
	fmt.Println("CPUs:",runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())


	var counter int64

	const gs = 100 // go routines
	var wg sync.WaitGroup //WAITGROUP
	wg.Add(gs)


	for i := 0; i < gs; i++ {
		go func(){
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()  // yeilds the processor(CPUs), allowing other Go routines to run, basically, you can run something else
			fmt.Println("Counter baby:", atomic.LoadInt64(&counter))
			wg.Done() // WaitGroup -this has to be in your go func
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}

	wg.Wait() // WaitGroup - wait for all go func to be done
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count:", counter) 
}