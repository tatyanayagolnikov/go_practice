package main

import "fmt"
import "runtime"
import "sync"

/// RACE CONDITION ///


func main() {
	fmt.Println("CPUs:",runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())


	counter := 0

	const gs = 100 // go routines
	var wg sync.WaitGroup //WAITGROUP
	wg.Add(gs)

	var mu sync.Mutex // MUTEX

	for i := 0; i < gs; i++ {
		go func(){
			mu.Lock() // Mutex 
			v := counter 
			runtime.Gosched()  // yeilds the processor(CPUs), allowing other Go routines to run, basically, you can run something else
			v++
			counter = v 
			mu.Unlock() // Mutex
			wg.Done() // WaitGroup -this has to be in your go func
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}

	wg.Wait() // WaitGroup - wait for all go func to be done
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count:", counter) 
}