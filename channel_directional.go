package main

// CHANNELS BLOCK! 

import "fmt"

func main() {

	c := make(chan int) // make a CHANNEL

	cr := make(<-chan int) // RECEIVE CHANNEL

	cs := make(chan<- int) // SEND CHANNEL 
	
	fmt.Println("----------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// general to specific assignments 
	cr = c
	cs = c
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
	
}
