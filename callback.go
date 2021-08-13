package main

/// a CALLBACK is Passing a Func as an Argument ////

import "fmt"

func main() {
	// ii := []int{1,2,3,4,5,6,7,8,9}
	// s := sum(ii...)
	// fmt.Println(s)
	

	// ii2 := []string{"sweet", "tweet",}
	// s2 := choc(ii2...)
	// fmt.Println(s2)
    
	// b2 := []int{2,4,6,8}
	// bb := bunk(b2...)
	// fmt.Println(bb)
	// fmt.Println(b2)
    
	// b3 := []int{9,8,7,6,5,4,3,2,1}
	// bbb := bunk(b3...)
	// fmt.Println(bbb)

	// b4 := []string{"beets", "bears", "battlestar galactica", "smoikin"}
	// bbbb := choc(b4...)
	// fmt.Println(bbbb)
	// fmt.Println(b4)

	// iii := []string{"tanya", "the", "programmer",}
	// ss := sum2(iii...)
	// fmt.Println(ss)

	// b6 := []int{9,8,7,6,5,4,3,2,1}
	// fmt.Println(b6)
	// sss := sum(b6...)
	// fmt.Println(sss)

	// st := even(sum, ii...)
	// fmt.Println("even numbers:", st)

	// gg := odd(bunk, b6...)
	// fmt.Println("the odd numbers:", gg)

	// rr := yeah(sum2, b4...)
	// fmt.Println("it works!", rr)
	
	// i5 := []string{"we will", "we will", "rock you",}
	// b8 := boo(i5...)
	// fmt.Println(b8, i5)
    xs := []string{"bears", "beets", "battlestar galactica"}
	f := foo(xs...)
	fmt.Println(f, xs)


}


func foo(m ...string) string {
	return "smells good"
}

// func boo(b ...string) string {
// 	return "bees and honey"
// }

// func yeah(n func(xi ...string) string, ss ...string) string {
// 	return "something"

// }

// func sum2(xi ...string) string {
// 	return "smurt"
// }

// func odd(bb func(b ...int) int,  vi ...int) int {
// 	var dd []int
// 	for _, v := range vi {
// 		if v % 2 != 0 {
// 			dd = append(dd, v)
// 		}
// 	}
// 	return bb(dd...)
// }

// func even(f func(xi ...int) int, vi ...int) int {
// 	var yi []int
// 	for _, v := range vi {
// 		if v % 2 == 0 {
// 			yi = append(yi, v)
// 		}
// 	}
// 	return f(yi...)
// }

// func bunk(b ...int) int {
// 	total := 0
// 	for _, v := range b {
// 		total += v 
// 	}
// 	return total 
// }



// func choc(l ...string) string {
// 	return "mission" 
// } 


// func sum(xi ...int) int {
// 	total := 0
// 	for _, v := range xi {
// 		total += v 
// 	}
// 	return total 

// }




