package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle 
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}
type person struct {
	first string
	last string
	favflav []string
}

func main() {
	pp := person {
		first: "snale",
		last: "whale",
		favflav: []string{
			"water",
			"cola",
			"cream soda",
		},
	}
////// MAP - map[string]int 
//////// MAPS - KEY & VALUE store - super fast and efficient lookup - allows you to store some values based upon some key
/////// var m := "map" "[KEY = "type"]" VALUE = "type" and then a composite literal "{}"
/////// m := map[string]int{} ///// COMPOSITE LITERAL IS THE TYPE "map[string]int" is the type in the example on the left
	m := map[string]person{
		pp.last: pp, 
		pp.first: pp,
	}
	

	for _, val := range m{
		fmt.Println(val.first)
		for i, v := range pp.favflav{
			fmt.Println(i, v)
		}
	}


	// for i, v := range pp.favflav{
	// 	fmt.Println(i, v)
	// }



//// ANONYMOUS STRUCT ////  s := struct {fields} {values}
	p1 := struct {
		first string
		friends map[string]int
		fav []string
	}{
		first: "stanly",
		friends: map[string]int {
			"chrees": 206,
			"tatbot": 425,
		},
		fav: []string{
			"water",
			"watermelon",

		},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.friends)

	for k, v := range p1.friends {
		fmt.Println(k, v)
	}
	for i, val := range p1.fav {
		fmt.Println(i, val)
	}

	
	// t := truck {
	// 	fourWheel: true,
	// 	vehicle: vehicle{
	// 		doors: 4,
	// 		color: "black",
	// 	},
	// }
	// s := sedan {
	// 	luxury: true,
	// 	vehicle: vehicle{
	// 		doors: 4,
	// 		color: "grey",
	// 	},
	// }
	// fmt.Println(t)
	// fmt.Println(t.doors)
	// fmt.Println(t.color)

	// fmt.Println(s)
	// fmt.Println(s.doors)
	// fmt.Println(s.color)
}