package main

import (
	"fmt"
)

func main() {
	foo()

	bar(22)

	s1 := woo("whoopie 3")
	fmt.Println(s1)

	x, y := tang("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

	h, i := gong("Smoth", 33)
	fmt.Println(h)
	fmt.Println(i)
	
}
func gong(fn string, g int) (string, string) {
	a := fmt.Sprint(fn, ", and I am ", g, "years")
	b := fmt.Sprint("pheww")
	return a, b
}
func tang(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ` says "Hello`)
	b := false
	return a, b

}
func foo() {
	fmt.Println("hi from foo")
}
func bar(s int) {
	fmt.Println("hello bar", s)
}
// RETURN //
func woo(s string) string {
	return fmt.Sprint("Hey from, ", s)

}


// FUNCTIONS Syntax - func (r reciever) identifier(parameters(s)) (return(s)) { code ... }
