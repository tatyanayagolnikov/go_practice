package main

import "fmt" 
//// EXERCISE #5 ///////
// import "math" 
       

func main() {	
////// EXERCISE #1 ///////
// f := foo()
// fmt.Println(f)	

// b, a := bar()
// fmt.Println(b, a)

////// EXERCISE #2 ///////
// xi := []int{1,2,3,4}
// fmt.Println(xi)
// f := foo(xi...)
// fmt.Println(f)

// xii := []int{3,4,5}
// x2 := bar(xii)
// fmt.Println(x2)

//// EXERCISE #3 ///////
// defer foo()
// defer bar()


//// EXERCISE #4 ///////
// p1 := person {
// 	first: "Tanya",
// 	last: "Programmer",
// 	age: 35,
// }
// p1.speak()

//// EXERCISE #5 ///////
// c := circle {
// 	radius: 12.345,
// }

// s := square {
// 	length: 15,
// }

// info(c)
// info(s)

//// EXERCISE #6 ///////
// func(x int, y string) {
// 	fmt.Println("this is an anonymous func")
// }(42,"yes")

// func(s string, y int) {
// 	fmt.Println("this is another anonymous func")
// }("beets",55)

// func() {
// 	for i := 0; i < 100; i++ {
// 		fmt.Println(i)
// 	}
// }()

//// EXERCISE #7 ///////
// f := func() {
// 	fmt.Println("this is a func assigned to a variable")
// }

// x := func() {
// 	fmt.Println("x func assigned to variable")
// }

// f()
// x()
// g = f 
// g()
// fmt.Printf("this is g %T\n", g)

//// EXERCISE #8 ///////
// f := foo()
// fmt.Println(foo()())

//// EXERCISE #9 ///////
// x := []int{1,2,3,4,5,}
// y := sum(x...)
// fmt.Println("the total sum of all numbers is", y)

// e := even(sum, x...)
// fmt.Println("Even numbers", e)

// o := odd(sum, x...)
// fmt.Println("Odd numbers", o)


//// EXERCISE #10 ///////
// f := num()
// fmt.Println("this works", f)

// f2 := fun()
// fmt.Println("alrighty then", f2)

//// EXERCISE #9-2 ///////
// g := func(xi [] int) int {
// 	return xi[0] + xi[len(xi)-1]
// }

// x := foo(g, []int{1,2,3,4,5,6})
// fmt.Println(x)






// f := foo()
// fmt.Println(f)

// }


// func foo() func() int {
// 	x := 0
// 	return func() int {
// 		return x 
// 	}
	
f := foo()
fmt.Println(f())
fmt.Println(f())

}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}



//// EXERCISE #9-2 ///////
// func foo(f func(xi [] int)int, ii []int) int {
// 	n := f(ii)
// 	n++
// 	return n 
// }

//// EXERCISE #10 ///////
// func num() func () int {
// 	x := 0 
// 	return func () int {
// 		x++
// 		return x 
// 	}
// }

// func fun() func () int {
// 	var x int
// 	return func() int {
// 		x--
// 		return x
// 	}
// }

//// EXERCISE #9 ///////
// func sum(xi ...int) int {
// 	total := 0 
// 	for _, v := range xi {
// 		total += v
// 	}
// 	return total
// }

// func even(f func(xi ...int) int, vi ...int) int {
// 	var yi [] int
// 	for _, v := range vi {
// 		if v % 2 == 0 {
// 			yi = append(yi, v)
// 		}
// 	}
// 	return f(yi...)
// }

// func odd(f func(xi ...int) int, vi ...int) int {
// 	var yi [] int
// 	for _, v := range vi {
// 		if v % 2 != 0 {
// 			yi = append(yi, v)
// 		}
// 	}
// 	return f(yi...)
// }


//// EXERCISE #8 ///////
// func foo() func() int {
// 	return func() int{
// 		return 42
// 	}
// }

//// EXERCISE #7 ///////
// var g func()

//// EXERCISE #5 ///////
// type circle struct {
// 	radius float64
// }
// type square struct {	
// 	length float64
// }
// type shape interface {
// 	area() float64 
// }
// func (c circle) area() float64 {
// 	 return math.Pi * c.radius * c.radius 
	
// }
// func (s square) area() float64 {
// 	 return s.length * s.length 
	
// }
// func info(s shape) {
// 	fmt.Println(s.area())
// }



//// EXERCISE #4 ///////
// type person struct {
// 	first string
//     last string
//     age int
// }

// func (p person) speak() {
// 	 fmt.Println("I am", p.first, "the", p.last, "and I am", p.age, "years old")
// }


//// EXERCISE #3 ///////
// func foo() {
// 	func() {
// 		defer fmt.Println("defered anonymous func")
// 	}() 
// 	fmt.Println("yes")
// }

// func bar() {
// 	 fmt.Println("okay")
// }


//// EXERCISE #2 ///////
// func bar(i []int) int {
// 	sum := 0
// 	for _, v := range i {
// 		sum += v
// 	}
// 	return sum
// }

// func foo(x ...int) int {
// 	sum := 0
// 	for _, v := range x {
// 		sum += v	
// 	}
// 	return sum 	
// }


////// EXERCISE #1 ///////
// func bar() (int, string) {	
// 	return 42, "big dawg"	
// }

