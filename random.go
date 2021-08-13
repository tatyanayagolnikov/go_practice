package main

import "fmt"

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

func main() {
/// ANONYMOUS STRUCT /// s := struct{feilds}{values}
	bn := struct{
		first string
		age int
		foods []string // slice of type string
		numbers []int // slice of type int
		sports map[int]string //map with key of type int, value of type string
	}{
		sports: map[int]string{
			1: "bball",
			2: "snow",
			3: "hiking",
		},
		first: "json",
		age: 37,
		foods: []string{"sushi", "RR", "ipa",},
		numbers: []int{1,2,3},
	}
	fmt.Println(bn.numbers)
	fmt.Println(bn)

	for k, v := range bn.numbers{
		fmt.Println(k, v)
	}
	for i, val := range bn.foods{
		fmt.Println(i, val)
	}


	t := truck {
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}
	fmt.Println(t.color)
	s := sedan {
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(s.doors)
}