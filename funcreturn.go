package main
// func (r reciever) identifier(p parameters) (return(s)) {code}
import "fmt"

/// RETURN A FUNC FROM A FUNC ///

func main () {
	
	s1 := foo()
	fmt.Println(s1)
	
	x := bar()
	fmt.Printf("%T\n", x)
	i := x()
	fmt.Println(i)

	// x1 := pop()
	// fmt.Printf("%T\n", x1)
    // ii := x1()
	fmt.Println(pop()()) // <-- cleaned up version

	x2 := top()
	 z := x2()
	 fmt.Println(z)
	 // fmt.Println(top()()) // <-- cleaned up version


}

func foo() string {
	return "hello world"
	
}

func bar() func() int {
	return func() int{
		return 4242
	}
}

func pop() func() string {
	return func() string {
		return "luke this is your monkey"
	}
}

func top() func() string {
	return func() string {
		return "this makes sense"
	}
}