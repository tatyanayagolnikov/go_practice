package main

import "fmt"
 func main() {
	 c := make(chan int)
	 go func() {
		 c <- 42 
	 }()

	 v, ok := <-c
	 fmt.Println(v, ok)
	 
	
	
	fmt.Println("Tanya Programmer")
 }











































// package main 

// import "fmt" 

// func main() {
// 	even := make(chan int)
// 	odd := make(chan int)
// 	quit:= make(chan bool)

// 	// Send
// 	go send(even, odd, quit)
// 	// Receive
// 	receive(even, odd, quit)
	
	
	
// 	fmt.Println("Tanya Programmer")
// }

// // Send
// func send(even, odd chan<- int, quit chan<- bool) {
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			even <- i
// 		} else {
// 			odd <- i
// 		}
// 	}
// 	close(quit)
// }
// // Receive
// func receive(even, odd <-chan int, quit <-chan bool) {
// 	for {
// 		select {
// 		case v := <-even:
// 			fmt.Println("even", v)
// 		case v := <-odd:
// 			fmt.Println("odd", v)
// 		case i, ok := <-quit:
// 			if !ok {
// 				fmt.Println("not okay", i, ok)
// 				return
// 			} else {
// 				fmt.Println("okay", i)
// 			}
// 		}
// 	}
// }