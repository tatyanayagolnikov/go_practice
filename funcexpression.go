package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("My first func expression")
	}
	f()
	p := func(x int, y string) {
		fmt.Println("oh yes baby", x, y)

	}
	p(100, "points for me!")

}