package main

import "fmt"
// key word - identifier - type 

// func (r reciever) identifier(parameters) return(s) {... code} 

type person struct {
	first string
	last string
}
type secretAgent struct {
	person 
	ltk bool
}
type candy struct {
	person
	age int
	cool bool
}

func (x person) bong() {
    fmt.Println("This is I, bong, and my name is ", x.first)
}
func (x candy) bong() {
    fmt.Println("This is I, candy bong, and my name is ", x.last)
}

func (p person) goo() {
	fmt.Println("yoyoyo I am person ", p)
}

func (t candy) goo() {
	fmt.Println("yoyoyo I am candy ", t)
}

func (z secretAgent) goo() {
	fmt.Println("yoyoyo I am secretAgent ", z)
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, s.ltk, "the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "the person speak")
}
 
type human interface {
	speak()
	
}
type human2 interface {
	goo()
}


func jars (t human2) {
	fmt.Println("I am interfaces ", t)
}

func bar (h human) {
	fmt.Println("passed into bar",h)

}

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
		last: "yes",
	}
	p2 := candy{
		person: person{
			"Snow",
			"Diggy",
		},
		age: 33,
		cool: true,
	}


fmt.Println(sa1)
fmt.Println(sa2)
fmt.Println(p1)
fmt.Println(p2)
sa1.speak()
sa2.speak()
p1.bong()
sa1.bong()
sa2.bong()
p2.bong()

p2.goo()

bar(sa1)
bar(sa2)
bar(p1)
bar(p2)

jars(p1)
jars(p2)
jars(sa2)
jars(sa1)








}
