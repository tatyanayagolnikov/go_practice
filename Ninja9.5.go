package main

import "fmt"
//import "runtime"
import "sync"
import "sync/atomic"

// MUTEX - Mutual Excllution lock


func main() {
	var wg sync.WaitGroup // WAITGROUPS

	var incrementor int64

    gr := 100 // GO ROUTINES

	wg.Add(gr) // WAITGROUP - adds delta to counter


	for i := 0; i < gr; i++ {
		go func () {
			atomic.AddInt64(&incrementor, 1)

			fmt.Println("Middle value from func:", atomic.LoadInt64(&incrementor))
			
			wg.Done() // WAITGROUP - decrements counter by one(1)
		}()
	}
	wg.Wait() // WAITGROUP - blocks until counter is zero(0)
fmt.Println("End Value:",incrementor)

}