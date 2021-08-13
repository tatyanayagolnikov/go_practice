package main

import "fmt"
 
func main() {
	// x := 0 
	// fmt.Println("x befor",&x)
	// fmt.Println("x befor",x)
	// foo(&x)

	// fmt.Println("x after",&x)
	// fmt.Println("x after",x)


// y := 0
// fmt.Println(y)
// dog(&y)
// fmt.Println(y)

// z := 5
// fmt.Println(z)
// koo(&z)
// fmt.Println(z)

d := 98
s := 99
foo(&s)
foo(&d)

fmt.Println("1",s)
fmt.Println("2",d)

q := 25
got(&q)
fmt.Println(q)

}

func foo(x *int) {
	fmt.Println("first foo",x)
	fmt.Println("first foo",*x)
	*x = 42 
	fmt.Println("second foo",x)
	fmt.Println("second foo",*x)
}

func got(i *int) {
	fmt.Println("got 1",i)
	fmt.Println("got 1 value",*i)
	*i = 42
	fmt.Println("got 2",i)
	fmt.Println("got 2 value",*i)
	var g *int = i 
	fmt.Println("this is g",g)
}





// func koo(x *int) {
// 	*x = 1001
// 	fmt.Println(x)
// }

// func dog(x *int) {
// 	fmt.Println("func dog one",x)
// 	fmt.Println("func dog one",*x)
// 	*x = 45
// 	fmt.Println("func dog two",x)
// 	fmt.Println("func dog one",*x)
// }


// func foo(y *int) {
// 	fmt.Println("y befor",y)
// 	fmt.Println("y befor",*y)
// 	*y = 43
// 	fmt.Println("y after",y)
// 	fmt.Println("y after",*y)
// }