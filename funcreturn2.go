package main

import "fmt"

func main() {

	fmt.Println(foo())

	x := bar()
	fmt.Printf("%T\n", x)
	fmt.Println(x()) // calling x

	y := nar()
	z := y() // assigned y() to variable z 
	fmt.Println(z)
	fmt.Printf("%T\n", y)

	/// cleaned up return for nar() ///
	// fmt.Println(nar()())


	}
	



func foo() string {
	return "Hello world"
}

func bar() func() int {
	return func() int{
		return 42
	}
}

func nar() func() string {
	return func() string {
		return "WOOHOO"
	}
}