package main

import "fmt"

func main() {
	// var v int = 42
	// fmt.Println(v)
	// fmt.Println(&v)

	p1 := person {
		name: "tanya programmer",
		sport: "basketball",

	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
	



}

type person struct {
	name string
	sport string 
}

func changeMe(p *person) {
	p.name = "Tanya THE Programmer"
	// (*p).name = "Tanya is THE PROGRAMMER"
	p.sport = "bball baby!"
	// (*p).sport = "Baller"
}
