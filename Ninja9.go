package main

import "fmt"
import "runtime"
import "sync"

var wg sync.WaitGroup

func main() {
	fmt.Println("Begin: CPUs\t", runtime.NumCPU())
	fmt.Println("Begin: Goroutines\t", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go sho()
	fmt.Println("Middle: CPUs\t", runtime.NumCPU())
	fmt.Println("Middle: Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("End: CPUs\t", runtime.NumCPU())
	fmt.Println("End: Goroutines\t", runtime.NumGoroutine())

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Tanya Programmer", i)
	}
	wg.Done()
}

func sho() {
	for i := 5; i < 10; i++ {
		fmt.Println("Tat the PG", i)
	}
	wg.Done()
}