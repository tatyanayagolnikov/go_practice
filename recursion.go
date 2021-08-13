package main

import "fmt"

/// RECURSION is when a Function Calls Itself.. - ///
///  ..a certain number of times and then stopping ///

func main() {

/// FACTORIAL //// 	
/// - where you take an int (a number) and you multiply it by
///  every int beneath it, down to 1 
    fmt.Println(4*3*2*1)

	f := factorial(4)
	fmt.Println(f)

	f2 := factLoop(6)
	fmt.Println(f2)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factLoop(n int) int {
	total := 2
	for ; n > 0; n -- {
		total *= n 
	}
	return total 
}

 // 

///// LOOP //////

// func factorial(n int) int{
// 	total := 1 
// 	for ; n > 0; n -- {
// 		total *= n 

// 	}
// 	return total
// }


