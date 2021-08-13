package main

import "fmt"
//import "runtime"
import "sync"
//import "sync/atomic"

// MUTEX - Mutual Excllution lock


func main() {
	var wg sync.WaitGroup // WAITGROUPS

	incrementor := 0 // variable holding an incrementer value

    gr := 100 // GO ROUTINES

	wg.Add(gr) // WAITGROUP - adds delta to counter

	var mu sync.Mutex // MUTEX  



	for i := 0; i < gr; i++ {
		go func () {

			mu.Lock() // MUTEX

			v := incrementor // new variable [v] that reads and stores the incrementor variable value
			v++ // increments the new [v] variable 
			incrementor = v // writes the value stored in the new variable [v] back to incrementor variable
			fmt.Println(incrementor)
			
			mu.Unlock() // MUTEX
			wg.Done() // WAITGROUP - decrements counter by one(1)
		}()
	}
	wg.Wait() // WAITGROUP - blocks until counter is zero(0)
fmt.Println("End Value:",incrementor)

}