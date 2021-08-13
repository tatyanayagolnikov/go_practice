package main

import "fmt"

type hotdog int

func main() {

	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x) // conversion 
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}