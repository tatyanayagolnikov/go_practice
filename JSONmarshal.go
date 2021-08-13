package main

//// MARSHAL - //// (send data to someone using Marshal)
// If you want to take Data, and turn it into..
//..JSON, right in your program and assign it to..
//..a vairiable, you're going to MARSHAL  it.
import "fmt"
import "encoding/json"


type person struct {
	First string  // Fields are uppercase to marshal JSON
	Last string 
	Age int
}

func main() {
	
	p1 := person {
		First: "Tanya",
		Last: "Programmer",
		Age: 35,
	}
	p2 := person {
		First: "fuzzy",
		Last: "creatures",
		Age: 1,
	}
	people := []person {
		p1,
		p2,
	}

	fmt.Println(people)

	sb, err := json.Marshal(people) // MARSHAL 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb)) 


	

	// bs, err := json.Marshal(people)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(bs))
}
