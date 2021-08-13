package main

import "fmt"
import "math"

// Method sets are all about determining the interface the type implements

func main() {

c := circle {
	radius: 12.345,
}

info(&c) // info(&c)  <--this value is a pointer now, and will still run 

s := square {
	length: 15,
}
info(s)	
	
}

type circle struct {
	radius float64
}
type square struct {	
	length float64
}
type shape interface {
	area() float64 
}
func (c *circle) area() float64 {
	 return math.Pi * c.radius * c.radius 
	
}
func (s square) area() float64 {
	 return s.length * s.length 
	
}
func info(s shape) {
	fmt.Println(s.area())
}