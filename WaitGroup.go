package main

import "fmt"
import "runtime"
import "sync"

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS )
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
//	wg.Add(1)
	go foo()
	bar()
	go bee()
    wg.Add(2)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

}

func bee() {
	fmt.Println("Tanya Programmer")
	wg.Done()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}