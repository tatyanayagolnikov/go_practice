package main

import "fmt"
import "os"
import "io"

func main() {
	fmt.Println()
	fmt.Fprintln(os.Stdout, "Hello Tanya Programmer")
	io.WriteString(os.Stdout, "Tanya eez dee programmer")

}