package main

/// UNMARSHAL - ///
// If you have data coming into your program, 
// and you need to turn it into go data structure,
// (put it all in a go data structure), you're 
// going to UNMARSHAL it. 

import "fmt"
import "encoding/json"


func main() {
	s := `[{"First":"Tanya","Last":"Programmer","Age":35},{"First":"fuzzy","Last":"creatures","Age":1}]`
	bs := []byte(s)
	fmt.Println(bs)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	 people := []person {}
	// var people []person 

	err := json.Unmarshal(bs, &people) // second argument is an address of some type
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\n PERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}