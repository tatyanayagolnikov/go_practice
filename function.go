package main
//// everythilng in GO is PASS BY VALUE ////
import "fmt"

func main() {
// FUNCTIONS Syntax - func (r reciever) identifier(parameters(s)) (return(s)) { code ... }


// smar("shmee")
// tang()
// s2 := doo("your daddy")
// fmt.Println(s2)
// t, y := blang("blazer", "GGG")
// fmt.Println(t)
// fmt.Println(y)
// foo() 
// bar("James")
// s1 := woo("bond")
// fmt.Println(s1)
// x, y := mouse("Ian", "Fleming")
// fmt.Println(x)
// fmt.Println(y)


x := sum("james", 4,5,6)
fmt.Println("The total is, ", x)

y := sum("sarah", 22,33,55,68)
fmt.Println("this is different number", y)



}

func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0 
	for i, v := range x {
		sum += v 
		fmt.Println("for index position:", i, " we are adding ", v, " to the total sum ", sum )
	}
	fmt.Println("THE TOTAL IS: ", sum)
	return sum 
}

























// func blang(fn, ln string) (string, bool) {
// 	a := fmt.Sprint(fn, ln, `, say"Hello"`)
// 	b := true 
// 	return a, b 

// }

// func doo(st string) string {
// 	return fmt.Sprint("Luke this is,", st)

// }
// func smar(t string) {
// 	fmt.Println("dast is boot,", t)
// }
// func tang() {
// 	fmt.Println("tanya programmer")
// }
// func foo() {
// 	fmt.Println("yello foo")
// } 

// func bar(s string) {
// 	fmt.Println("yoyoyoy,", s)
// }
// func woo(s string) string {
// 	return fmt.Sprint("Hello form the woo, ", s)
// }
// func mouse (fn string, ln string) (string, bool) {
// 	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
// 	b := false
// 	return a, b
// }