package main

import "fmt"

// func (r reciever) identifier(p parameters) (return(s)) {code}

type person struct {
	name string
	sport string 
	age int 
}

func changeMe(p *person) {
	(*p).name = "Tanya the PROGRAMMER"
	(*p).name = "Tatertot the programmer"

	
}


func main() {
	p1 := person {
		name: "Tat Programmer",
		sport: "basketball baby",
		age: 35,
	}
	

	changeMe(&p1)
	p1.age = 42
	fmt.Println(p1)
	p1.name = "yes"
	//changeMe(&p1)
	fmt.Println(p1.name)

}