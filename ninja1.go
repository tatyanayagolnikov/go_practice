package main

import "fmt"



// var x int
// var y string
// var z bool

// var x = 42
// var y = "Jamesbond"
// var z = true

// type hotdog int 
// var x hotdog 
// var y int

// var x bool
// var x int
// var y float64


// var x int8 = 127

// const (
//     a int = 42
//     b float64 = 42.78
//     c string = "James Bond"
// )
// var a  = 42 >= 4
// var b  = (4 == 2) 
// var c  = (4 != 5)
// var d  = (5 < 6)
// var e  = (5 > 4) 
// var f  = (5 <= 5)
// const (
// 	a int = 2
// 	b bool = true
// 	c string = "mary had a little lamb"
// 	d = 4
// 	e = ""
// 	f = true 
// )

// const (
// 	a = 2021 + iota
// 	b 
// 	c 
// 	d 
// )

// const (
// 	a = iota -5
// 	b 
// 	c 
// 	d = iota 
// 	e 
// 	f 
// )



////////FOR LOOP////////////
// [for] init :=  ; condition ; post  {}
// "for" [condition | ForClause | Range Clause] Block {}


// func main() {
// 	for i := 0; i <= 10; i++ {
// 		for j := 0; j < 3; j++ {
// 			fmt.Printf("The outer loopp: %d\t The inner loop: %d\n", i, j)
// 		}
// 	}



/////////// STRUCT - COMPOSITE DATA STRUCTURE - AGGREGATE DATA STRUCTURE - COMPLEX DATA STRUCTURE ////// "type" "any name" "struct" {put in some fields, ex: first string}
// type person struct {
// 	first string
// 	last string
// 	age int
// }

// type secretAgent struct {
// 	person
// 	ltk bool
// }

// type hotdog struct {
// 	person
// 	secretAgent
// 	style string
// }

/// we create VALUES of a certain TYPE that are stored in VARIABLES
/// and those VARIABLES have IDENTIFIERS

// type person struct {
// 	first string
// 	last string
// 	fav_icecream []string
// }
type vehicle struct {
	doors string
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
	t := truck {
		fourWheel: true,
		vehicle: vehicle{
			doors: "four door",
			color: "black",
		},
	}
	fmt.Println(t.doors)
	fmt.Println(t.color)
	fmt.Println(t)
	fmt.Println(t.fourWheel)


	// x := person {
	// 	first: "peter",
	// 	last: "dude",
	// 	fav_icecream: []string{
	// 		"vanilla", 
	// 		"choc", 
	// 		"mint",
	// 	},
	// }
	// y := person {
	// 	first: "Stanly",
	// 	last: "bunny",
	// 	fav_icecream: []string{
	// 		"carrot",
	// 		"celery",
	// 		"apple",
	// 	},
	// }

	// m:= map[string]person{
	// 	x.last : x,
	// 	y.last : y,
	// }
	// fmt.Println(m)
	
	// for _, v := range m {	
	// 	fmt.Println(v.first)
	// 	fmt.Println(v.last)
	// 	for i, val := range v.fav_icecream {
	// 		fmt.Println(i, val)
	// 	}
	// 	fmt.Println("-----------")
	// }


	// fmt.Println(x.first, x.last)
	// for i, v := range x.fav_icecream {
	// 	fmt.Println(i, v)
	// }

	// fmt.Println(y.first, y.last)
	// for i, v := range y.fav_icecream {
	// 	fmt.Println(i, v)
	// }






/////// ANONYMOUS STRUCTS ////////	s := struct{fields}{}
	// p1 := struct {
	// 	first string
	// 	last string
	// 	age int
	// }{
	// 	first: "James",
	// 	last: "bond",
	// 	age: 32,
	// }
	// fmt.Println(p1)



/////////// STRUCT - COMPOSITE DATA STRUCTURE - AGGREGATE DATA STRUCTURE - COMPLEX DATA STRUCTURE //////
	// f2 := hotdog{
	// 	person: person{
	// 		first: "Costco",
	// 		last: "Foods",
	// 		age: 22,
	// 	},
	// 	secretAgent: secretAgent{
	// 		ltk: true,
	// 	},
	// 	style: "spicey",
	// }
	// sa1 := secretAgent{
	// 	person: person{
	// 		first: "James",
	// 		last: "Bond",
	// 		age: 32,
	// 	},
	// 	ltk: true,
	// }
	// // p1 := person{
	// // 	first: "James",
	// // 	last: "Bond",
	// // 	age: 32,
	// // }
	// p2 := person{
	// 	first: "Miss",
	// 	last: "Moneypenny",
	// 	age: 27,
	// }
	// fmt.Println(f2.age)
	// fmt.Println(sa1)
	// fmt.Println(p2)
	// fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	// fmt.Println( p2.first, p2.last, p2.age)



	////////// range ////////////

	//  x := [5]int{1,2,3,4,5}
	// fmt.Println(x)

	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T\n", x)

	// x := []int{1,2,3,4,5,6,7,8,9,22}
	// fmt.Println(x)

	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }

	// fmt.Printf("%T\n", x)


	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// fmt.Println(x)
	// fmt.Println(x[:5])
	// fmt.Println(x[5:])
	// fmt.Println(x[2:7])
	// fmt.Println(x[1:6])

////////// APPEND ////////////////	
	// x := []int{42,43,44,45,46,47,48,49,50,51}
	// x = append(x,52)
	// fmt.Println(x)
	// x = append(x, 53,54,55)
	// fmt.Println(x)
	// y := []int{56,57,58,59,60}
	// x = append(x,y...)
	// fmt.Println(x)

	// x := []int{42,43,44,45,46,47,48,49,50,51}
	// fmt.Println(x)
	// x = append (x[:3], x[6:]...)
	// fmt.Println(x)
	



/////// COMPOSITE LITERAL
/////// x := type{value} // COMPOSITE LITERAL
/////// a SLICE allows you to group together VALUES of the same TYPE
	// x := []int{4,5,7,8,42} ////// x := slice of[], Type int { values 4, 5, 6, 7, 42 etc }
	// fmt.Println(x)
	// fmt.Println(len(x))

	//  x := []int{4,5,6,7,8,}

	//  fmt.Println(x[1])
	//  fmt.Println(x)
	//  fmt.Println(x[1:1])
	//  fmt.Println(x[2:3])

	//  for i, v := range x {
	// 	 fmt.Println(i, v)
	//  }

	//  for i := 0; i <= 4; i ++ {
	// 	 fmt.Println(i, x[i])
	//  }


///////////////  ... (dot dot dot)is a VARIADIC PARAMETER, which means an unlimited number of values of T(this type) can be passed in 
	// x := []int{1,2,3,4}
	// y := []int{5,6,7,8}
	// y = append(y,x...)
	// fmt.Println(y[3]) //prints Index Position 3 (index position starts at 0)
	// for i, v  := range x {
	// 	fmt.Println(v, i)
	// }
	

//////////////// APPEND ////////
	// x := []string{"one","two","three"}
	// y := []string{"Tanya","programmer"}
	// x = append(x,y...)
	// fmt.Println(x)

	// x := []int{1,2,3}
	// fmt.Println(x)
	// x = append(x, 11,22,33,44)
	// fmt.Println(x)

	// y := []int{123,223,333}
	// x = append(x,y...)
	// fmt.Println(x)

	// x = append(x[:2], x[5:]...)

	// fmt.Println(x)
	
	
	// x := make([]string, 50, 50)
	// x = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`, }
	// fmt.Println(cap(x))
	// fmt.Println(len(x))
 
	// for i := 0; i < len(x); i ++ {
	// 	fmt.Println(i, x[i])
	// }

// xs1 := []string{"James", "bond", "Shaken"}
// xs2 := []string{"Miss", "mooneh", "Heyaaa", "James" }
// fmt.Println(xs1)
// fmt.Println(xs2)

// xxs := [][]string{xs1, xs2}

// fmt.Println(xxs)

// for i, xs := range xxs {
// 	fmt.Println("record: ", i, xs)
// 	// for j, val := range xs {
// 	// 	fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
// 	// }
// }


/////// MAKE ///// specify how large the slice should be, and how large the underlaying slice should be
//////x:= make([]type, length, capacity)/////////

	// x := make([]int, 10, 12)
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))
	// x[0] = 42
	// x[9] = 99
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// x = append(x, 2222)
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// x = append(x, 2223)
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	// x = append(x, 2224)
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))


/////// MULTI-DIMENTIONAL SLICE - slice within a slice
	// jb := []string{"James", "Bond", "Choc", "Martini"}
	// fmt.Println(jb)
	// mp := []string{"Miss", "Penny", "Berry", "Gum"}
	// fmt.Println(mp)

	// xp := [][]string{jb, mp}
	// fmt.Println(xp)



// 	m := map[string][]string{
// 		"bond_james": []string{`Shaken, not stirred`, `Martinis`, `Women`},
// 		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
// 		"no_dr": []string{`Being evil`,`Ice cream`, `Sunsets`,},
// 	}
// /////// ADDING a Record (Key and Value) to a MAP /////
// 	m["big_poppa"] = []string{"bbqs", "cold brews", "petting zoos"}

// /////// DELETING a Record from a MAP ///////////
//     delete (m, "no_dr")

// 	for k, v := range m {
// 		fmt.Println("this is the record for: ", k)
// 		for i, v2 := range v {
// 			fmt.Printf("\t index position: %v \t value: %v \n", i, v2)
// 		}
// 	}
	



// 	x := map[string][]int{
// 		"thing1": []int{1,2,3},
// 		"thing2": []int{4,5,6},
// 		"thing3": []int{7,8,9},
// 	}
// 	fmt.Println(x)
// ///// ADD a Value to a Key /////////
// 	x["thing1"] = []int{55}

// 	for k, v := range x {
// 		fmt.Println("This is the record for: ", k)
// 		for i, v2 := range v {
// 			fmt.Println("\t The index and value are: ", i, v2)
// 		}
// 	}







////// MAP - map[string]int 
//////// MAPS - KEY & VALUE store - super fast and efficient lookup - allows you to store some values based upon some key
/////// var m := "map" "[KEY = "type"]" VALUE = "type" and then a composite literal "{}"
/////// m := map[string]int{} ///// COMPOSITE LITERAL IS THE TYPE "map[string]int" is the type in the example on the left
    


// m := map[string]int{
	// 	"James":32,
	// 	"Penny":27,
	// } 
	// fmt.Println(m)
	// fmt.Println(m["James"])

	// fmt.Println(m["Pearls"])
	// v, ok := m["Pearls"]
	// fmt.Println(v)
	// fmt.Println(ok)

	// if v, ok :=m["Pearls"]; ok {
	// 	fmt.Println("This is the if print", v)
	// }
	// if v, ok :=m["Penny"]; ok {
	// 	fmt.Println("This is the if print", v)
	// }	




	// m := map[string]int{
	// 	"a-ames": 32,
	// 	"b-Pennt": 27,
	// }
	// fmt.Println(m)

//////// DELETE - value from a slice /////////
	// delete (m, "a-ames")

	// delete (m, "spoon")
	// fmt.Println(m)

	// fmt.Println(m["b-Pennt"])
	// fmt.Println(m["spoon"])

	// if v, ok := m["spoon"]; ok {
	// 	fmt.Println("value", v)
	// 	delete(m, "a-ames")
	// } 
	// fmt.Println(m)



	// fmt.Println(m["James"])
	
	// fmt.Println(m["snow"])

	// v, ok :=m["snow"]

	// fmt.Println(v)
	// fmt.Println(ok)

////////////// ADD record to MAP /////////////

	// m["todd"] = 33
	// // m["tat"] = 35

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// xi := []int{1,2,3,4,5}

	// for i, v := range xi {
	// 	fmt.Println(i, v)
	// }


	// if v, ok := m["snow"]; ok {
	// 	fmt.Println("This is okay to print", v)
	// }

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// xi := []int{4,5,6,7,8,9}

	// for i, v := range xi {
	// 	fmt.Println(i, v)

	// }

	// delete(m, "snow")
	// fmt.Println(m)





	 


	

//////////for INDEX(i), VALUE(v) := range (var x) {}
	//  for i, v := range x { 
	// 	 fmt.Println(i, v)
	//  }
	



	// var x [5]int
	// fmt.Println(x)
	// x [3] = 42
	// fmt.Println(x)
	// fmt.Println(len(x))

///////////// FOR LOOP ////////////////	

	// z := 42
	// for z <= 42 {
	// 	fmt.Println(z)
	// }
 
	// for x := 65; x <= 90 ; x++ {
	// 	fmt.Println(x)
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("\t%#U\n", x)
	// 	}
	// 	// fmt.Printf("%#U\n%#U\n%#U\n", x, x,x)
	// }

	// for b := 1985; b <= 2021; b++ {
	// 	fmt.Println(b, "alive")
	// }

	// bd := 1985
	// for bd <= 2021 {
	// 	fmt.Println(bd)
	// 	bd ++
	// }

	// de := 42
	// for { 
	// 	if de <= 43 {
	// 		break
	// 	} 
	// 	fmt.Println(de)
	// 	de ++
	// }

	// for c := 10; c <= 100; c ++ {
	// 	{
	// 		fmt.Printf("%v divided by 4, remainder (aka modulus) is %v\n", c, c%4 )
	// 	}
	// }

//////////// IF - ELSE IF - ELSE /////////////

	//  d := 5 
	// 	if d == 4 {
	// 		fmt.Println("woop")
	// 	} else if d < 10 {
	// 		fmt.Println("whos your daddy")
	// 	} else {
	// 		fmt.Println("not today charlie")
	// 	}

//////////// SWITCH //////////////

	// 	switch {
	// 	case false:
	// 		fmt.Println("22")
	// 	case true: 
	// 		fmt.Println("42")
	// 	}
	
	// 	f := "favSport"
	// 	switch f  {
	// 	case "proop": 
	// 	fmt.Println("pogramming is sweet") 
	// 	}
	// 	switch "favSport" {
	// 	case "smort", "goop", "favSport": 
	// 	fmt.Println("who is dee smoriest?")
	// 	fallthrough
	// 	case "favSp":
	// 		fmt.Println("yeee okie")
	// 	default: 
	// 	fmt.Println("how now brown cow")
	// 	}
		 
	// 	fmt.Printf("%v\n",true && true)
	// 	fmt.Printf("%v\n", true && false)
	// 	fmt.Printf("%v\n", true || true)
	// 	fmt.Printf("%v\n", true || false)
	// 	fmt.Printf("%v\n", !true)
	

	// for x := 0; x < 10001 ; x++ {
	// 	if x !=10001 {
	// 		fmt.Println(x)
	// 	}
	// }
	// fmt.Printf("%v\n",true && true)
	// fmt.Printf("%v\n",true && false)
	// fmt.Println(true || true)
	// fmt.Println(true || false)



	
	// a := 3
	// switch  {
	// case a != 4 :
	// fmt.Println("okay")
	// }


// 	switch {
// 	case false:
//       fmt.Println("somwbodeeeeeee")
// 	case true:
// 		fmt.Println("this is boop")
// 	}
// switch {
// case false:
// 	fmt.Println("should print")
// case true:
// 	fmt.Println("should print")
// }


	// a := "Smurt"
	// switch a {
	// case "burt", "corp", "Smurt":
	// 	fmt.Println("Toop boop moop")
	// case "urt":
	// 	fmt.Println("Toonya is smurt")
	// case "kurt":
	// 	fmt.Println("this is 2 ")
	// case 2 ==8:
	// 	fmt.Println("this is flase ")
	// case 2 <=1:
	// 	fmt.Println("this is wooooo ")
	// case 4 ==4:
	// 	fmt.Println("this is 4 baby!")
    // 	fallthrough
	// default: 
	// fmt.Println("Default mesage")
		
	// }

	





	// for i := 0; i < 100; i ++ {
	// 	if i%4 == 3 {
	// 		fmt.Println(i )
	// 	}
	// }



	// x := 44
	// if x == 42 {
	// 	fmt.Println("ye")
	// } else if x >= 99 {
	// 	fmt.Println("no")
	// } else if x <= 42 {
	// 	fmt.Println("no go jo")
	// } else if x >= 5 {
	// 	fmt.Println("yasss")

	// } else {
	// 	fmt.Println("This is Piggy bank")
	// }

	
	
		// fmt.Println("Tanya programs") ; fmt.Println("okay")
	// }
	// if 5 == 6 {
	// 	fmt.Println("woop")
	// }
	// if !false {
	// 	fmt.Println("Tanya programmer")
	// }



	// x := 56 % 40
	// y := 56 / 40
	// fmt.Println(x, y)



    // x := 0
	// for {
	// 	x ++
	// 	if x > 40 {
	// 		break

	// 	}
	// 	if x %2 != 0 {
	// 		continue 
	// 	}
	// 	fmt.Println(x)
			
	// }
	


// for i := 0; i < 3; i++ {
// 	fmt.Println(i)
// }


// for a := 33; a < 123; a ++ {
// 	fmt.Println(a)
// 	fmt.Printf("%s\t%#x\t%#U\n", a, a, a)

// }


	// x := 1
	// for x < 10 {
	// 	fmt.Println(x)
	// 	x ++
	// }
	// fmt.Println("done.")







	// for i := 0; i <= 10; i++ {
	// 	fmt.Printf("The outer loopp: %d\n", i,)
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("\t\t\t The inner loop: %d\n", j)
	
	// 	}
	// }
	





// fmt.Println(a)
// fmt.Println(b)
// fmt.Println(c)
// fmt.Println(d)

	// a := 42
	// fmt.Printf("%d\t%b\t%#X\n", a, a, a)
	// b := a << 1
	// fmt.Printf("%d\t%b\t%#X\n", b, b, b)




// 4 == 2 
// 5 <= 5
// 5 >= 4
// 5 != 6
// 5 < 8
// 5 > 


	// a := 42
	// fmt.Println(a)
	// fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n", c)
	// s := "Hello tanya programmer"
	// fmt.Println(s)
	// s = "Hello everyone"
	// fmt.Println(s)
	// fmt.Printf("%T\n", s)
	// fmt.Println([]byte(s))

	// bs := []byte(s)
	// fmt.Println(bs)
	// fmt.Printf("%T\n", bs)

	// for i := 0; i < len(s); i ++ {
	// 	fmt.Printf("%#U ", s[i])
	// }

	// fmt.Println("")

	// for i, v := range s {
	// 	fmt.Printf("at index position %d we have hex %#x\n", i, v)
	// }



	// fmt.Println(x)
	// fmt.Printf("%T\n", x)



	// x = 42
	// y = 42.34534
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Printf("%T\n", x)
	// fmt.Printf("%T\n", y)
	














	// fmt.Println(x)
	// fmt.Printf("%T\n", x)

	// x = 42
	// fmt.Println(x)
	
	// y = int(x)
	// fmt.Println(y)
	// fmt.Printf("%T\n", y)
//  x := 42
//  y := "Jamesbond"
//  z := true

    // fmt.Println(x, y, z)


	
	// fmt.Println(x)
	// fmt.Printf("%T\n", x)

	// fmt.Println(y)
	// fmt.Printf("%T\n", y)

	// fmt.Println(z)
	// fmt.Printf("%T\n", z)

	// s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	// fmt.Println(s)

}