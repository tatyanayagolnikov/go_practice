package main
//------------------
// CHANNELS BLOCK! 
//------------------
import "fmt"

func main() {

	// CHANNEL - succesful channel #1. - inside a go func()

	c := make(chan int) // make a CHANNEL of ints 

	go func() {
		c <- 42 // use [<-] notation to put something in the CHANNEL
	}()

	fmt.Println(<-c)
	fmt.Println("----------")
	fmt.Printf("%T\n", c)

	// CHANNEL - succesful channel #2
	// BUFFER CHANNEL - a Channel will hold x amount of values, until it bloscks

	c2 := make(chan int, 1) // CHANNEL, and add BUFFER CHANNEL(s) (almost any number) 

	c2 <- 25

	fmt.Println(<-c2)

	c3 := make(chan string, 1) // BUFFER amount is 1
	c3 <- "Tanya Programmer"
	fmt.Println(<-c3)





}