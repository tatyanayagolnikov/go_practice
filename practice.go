package main

import "fmt"
import "math"

// var a int
// type hotdog int
// var b hotdog

type square struct {
	length float64
}
type circle struct {
	diameter float64
}
type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.diameter * c.diameter 
}

func (s square) area() float64 {
	return s.length * s.length 
}

func info(s shape) {
	fmt.Println(s.area())
}


func main() {

c := circle {
	diameter: 3.3,
}

s := square {
	length: 4.4,
}

info(c)
info(s)
	
	






	
}