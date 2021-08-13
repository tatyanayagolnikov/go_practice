package main

import "fmt"

type person struct {
	first string
	last string
}
type secretAgent struct {
	person 
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "the person speak")
}

type human interface {
	speak()
}

func bar (h human) {
	switch h.(type) {
	case person:
		fmt.Println("this is purrrson", h.(person).first)
	case secretAgent:
		fmt.Println("this is secretAgent, muahahah", h.(secretAgent).first) 	

	}
	fmt.Println("I called human, passed into func bar", h)
}

type hotdog int


func main() {
	sa1 := secretAgent{
		person: person{
			"James", 
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"miss",
			"moneypenny",
		},
		ltk: false,
	}

	p1 := person{
		first: "Dr",
		last: "Yes",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
	
	// CONVERSION 
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)





}