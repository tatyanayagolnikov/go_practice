package main

import "fmt"

func main() {
	x := 2
	y := 3
	fmt.Println(x, y)
	x, y = y, x   // swap x and y 
	fmt.Println(x, y)
}