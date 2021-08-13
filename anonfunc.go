package main

import "fmt"



func main() {
/// ANONYMOUS FUNC ////
/// func(){}() ////

	func() {
		fmt.Println("anon func ran")
	}()

	func(x int) {
		fmt.Println("The meaning of life", x)
	}(42)

	func(t int, y string) {
		fmt.Println("double trouble", t, y)
	}(33, "Smith")
	
	
	foo()
}


func foo() {
	fmt.Println("yo foo")
}