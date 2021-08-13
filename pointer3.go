package main

import "fmt"
import "math" 


// func (r reciever) identifier(p parameters) (return(s)) {code}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// Method sets are all about determining the interface the type implements

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle {
		radius: 5,
	}
	fmt.Printf("%T\n", &c)
	info(&c)
	fmt.Println(c.area())
	fmt.Printf("%T\n", c)
	
}