package main

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
	secretAgent
	person
	sport string
}

// INTERFACES & POLYMORPHISM //

type human interface {
	speak()
	goat()
}

func (h hotdog) goat() {
	fmt.Println("My favorite sport is", h.sport, "my anme is", h.first, h.last)
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (s secretAgent) coo() {
	fmt.Println("Do I have a ltk?", s.ltk,)
}

func (s secretAgent) sol() {
	fmt.Println("Luke, I am", s.first, s.last)
}

func main() {
	p4 := hotdog{
		sport: "basketball",
		person: person{
			"Tat",
			"Bot",
		},
	}
	p4.goat()

	p3:= secretAgent{
		person: person{
			"YOUR",
			"Father",
		},
	}
	p3.sol()

	p2:= secretAgent{
		ltk: true,
	}
	p2.coo()

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

	fmt.Println(sa1)
	sa1.speak()  // sa1 [value] has access to the method [func speak()]
	sa2.speak()  //.. by just chaining it with a dot. 

}