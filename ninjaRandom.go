package main

import "fmt"
import "sort"

type user struct {
	First string
	Last string
	Age int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }


func main() {
	u1 := user{
		First: "Tanya",
		Last: "Programmer",
		Age: 35,
		Sayings: []string{
			"Shake it baby",
			"Make it",
			"Take it", 
		},
	}

	u2 := user{
		First: "Serge",
		Last: "Bro",
		Age: 37,
		Sayings: []string{
			"Take it easy",
			"Good job",
			"Dude, you can do it",
		},
	}

	u3 := user{
		First: "VenBen",
		Last: "LilBro",
		Age: 32,
		Sayings: []string{
			"Daddy like,",
			"I pitty da foo,",
			"Adventure time,",
		},
	}

	users := []user{
		u1,
		u2,
		u3,
	}

	sort.Sort(ByAge(users))
	for _, v := range users {
		fmt.Println("\n", "I am", v.First, v.Last, "and my age is", v.Age)
		sort.Strings(v.Sayings)
		for _, s := range v.Sayings {
			fmt.Println("\t",s)
		}
	}
fmt.Println("\t ---------------")
	sort.Sort(ByLast(users))
	for _, v := range users {
		fmt.Println("\n", "I am", v.First, v.Last, "and my age is", v.Age)
		for _, s := range v.Sayings {
			fmt.Println("\t",s)
		}
	}

	

}