package main

import "fmt"

func main() { // control flow enters with 1 go routine
	
	c := make(chan int) // make a CHANNEL 

	b := make(chan int) 

	p := make(chan string)

	// send
	go tall(b)
	// receive
	bar(b)



	// send 
	go ball(p)
	// receive
	coo(p)
	
	
	// send
	go rad(p) // launching a GO ROUTINE

	// receive  
	coo(p) 
	
	
	//send
	go foo(b) // launching a GO ROUTINE

	// receive 
	bar(b) 


	fmt.Printf("-----------\n")

	

	// send
	go foo(c) // launching a GO ROUTINE 

	// receive
	bar(c)
    
	fmt.Println("about to exit", "-Tat programmer")
}

// SEND (pass) - this is a Send Only Channel now
func foo(c chan<- int) {
	c <- 42 // passing in a value
}

// RECEIVE - this is a Receive Only Channel now
func bar(c <-chan int) {
	fmt.Println(<-c) // pulling the value of [<-c] off
}

// RECEIVE Channel only
func coo(s <-chan string) {
	fmt.Println(<-s)
}

// SEND Channel Only 
func rad(c chan<- string) {
	c <- "Tanya, the PROGRAMMER"
}

// SEND Channel Only 
func ball(b chan<- string) {
	b <- "Tatbot software engineer"
}

// SEND Channel Only
func tall(t chan<- int) {
	t <- 32
}
