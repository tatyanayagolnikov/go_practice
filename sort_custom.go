package main

import "fmt"
import "sort" 

type Person struct {
	Name string // first and age are FEILDS 
	Age int 
}
	func (p Person) String() string {
		return fmt.Sprintf("%s: %d", p.Name, p.Age)
	}
	
	// ByAge implements sort.Interface for []Person based on
	// the Age field.
	type ByName []Person
	
	func (a ByName) Len() int           { return len(a) }
	func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
	func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }


func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q Tanya Programmer", 35}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)

	sort.Sort(ByName(people)) // conversion 

	fmt.Println(people )
	
}