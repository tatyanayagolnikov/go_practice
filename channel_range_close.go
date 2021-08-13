package main

import "fmt"
// ------------------------------------------------ 
// a RANGE CLAUSE Blocks Until the CHANNEL is Closed
// --------------------------------------------------

func main() {
	
	c := make(chan int) // make a CHANNEL
	b := make(chan int) 
	
	// send
	go foo(b)

	// receive 
	bar(b)
	
	// send
	go tat(c)

	// receive
	bar(c)

	fmt.Println("about to exit. -Tanya Programmer") 
}

// SEND (pass) - this is a Send Only Channel now
func foo(c chan<- int) {
	for i:= 0; i < 5; i++ { // FOR LOOP - using init; condition; post
		c <- i // passing in a value
	}
	close(c) // CLOSE CHANNEL
}

func tat(t chan<- int) {
	for i := 0; i < 10; i++ {
		t <- i
	}
	close(t) // CLOSE CHANNEL
}

// RECEIVE - this is a Receive Only Channel now
func bar(c <-chan int) {
	for v := range c {
		fmt.Println(v) // pulling the value of [<-c] off
	}
}