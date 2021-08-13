package main

import "fmt" 

func main () {
	d := shwee("kyle")
	fmt.Println(d)
	x, y := dream("blue", 55)
	fmt.Println(x)
	fmt.Println(y)
	
	
}
func dream(cl string, ht int) (string, int) {
	a := fmt.Sprint("this is coo ", cl )
	b := 27
	return a, b
}
func shwee(k string) string {
	return fmt.Sprint("heyo ", k )
}