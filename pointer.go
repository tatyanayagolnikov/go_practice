package main

/// POINTER - Change the Value of an Address ///

import "fmt"

/// & (ampersand, which gives you the address, where a value is stored)


func main() {
	a := 42
	fmt.Println(a)
	//fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a 
	fmt.Println(b)
	//fmt.Println(*b) // * gives you the value stored at an address when you have the address
	//fmt.Println(*&a) // * de-reference 

	*b = 43 
	fmt.Println(&a)
	fmt.Println(a)


	*b = 45
	fmt.Println(b)







	// a := 42
	// fmt.Println(a)
	// fmt.Println(&a)

    // var b = &a 
	// fmt.Println("sweet",b)

	// *b = 44

	// var c = b
	// fmt.Println("tanya",c)

	// var d = c
	// fmt.Println("programmer",*d)
	


}
