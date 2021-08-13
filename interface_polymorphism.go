package main

// INTERFACES & POLYMORPHISM // - a VALUE can be of MORE than 1 TYPE
// - If you have these methods, you're also my type -


/// FUNC SYNTAX ///
// func (r reciever) identifier(parameters) (return(s)) { code }
// (if you have multiple returns, they need to be inside parenthesis(). )    
import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

type hotdog struct {
	person 
}


func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "secretAgent speak()", s.ltk)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "person speak()",)
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("Im called human from bar()",h)
}
 
//// ASSERTION ////

// func bar(h human) {
// 	switch h.(type) {      // [assertion] [.(type)] is specail type of switch statement where you can switch on type
// 	case person:
// 		fmt.Println("Im called human from person barrrrrr()", h.(person).first)
// 	case secretAgent:
// 		fmt.Println("Im called human from secretAgent barrrrrr()", h.(secretAgent).first)
// 	}
// 	fmt.Println("Im called human from bar()",h)
// }

func main() {
	p2 := hotdog{
		person: person{
			"hotttt",
			"doggo",
		},
	}

	p1 := person {
		first: "Timyyyy",
		last: "durrr",
	}

	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Tanya",
			"Programmer",
		},
		ltk: true,
	}

	sa1.speak()  // sa1 [value] has access to the method [func speak()]
	sa2.speak()  //.. by just chaining it with a dot. 

	bar(sa1)
	bar(sa2)
	bar(p1)
	bar(p2) 
	fmt.Println(p1)

}