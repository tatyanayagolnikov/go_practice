package main

import "fmt"

func main() {
	
	p1 := person{
		name: "Tanya Programmer",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}

type person struct {
	name string 
}

func changeMe(p *person) {
	fmt.Println(p)
	(p).name = "Tat is a great programmer"
	fmt.Println(p)
}