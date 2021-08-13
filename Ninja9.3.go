package main

import "fmt"
import "runtime"
import "sync"
//import "sync/atomic"




func main() {
	incrementor := 0 // variable holding an incrementer value

	const gr = 100 // GO ROUTINES

	var wg sync.WaitGroup // WAITGROUPS
	wg.Add(gr) // adds delta to WAITGROUP counter

	for i := 0; i < gr; i++ {
		go func () {

			v := incrementor // new variable [v] that reads and stores the incrementor variable value
			runtime.Gosched() // runtime.Gosched() yields the processor 
			v++ // increments the new [v] variable 
			incrementor = v // writes the value stored in the new variable [v] back to incrementor variable
			wg.Done() // decrements WAITGROUP counter by one(1)
		}()
		fmt.Println(incrementor)
	}
	wg.Wait() // blocks until WAITGROUP counter is zero(0)
fmt.Println(incrementor)

}