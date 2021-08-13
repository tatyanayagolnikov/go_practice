package main

import "fmt"
// func (r reciever) identifier(parameters) return(s) {... code} 







func main() {
	////// Func Expression //////
	// s := func () {
	// 	fmt.Println("ysssss")
	// }
	// s()
	
	fmt.Println("okay")
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("%T\n", x)
	i := x()
	fmt.Println(i)

	sam()
	fmt.Println(sam()())

	t := temp()
	c := t()
	fmt.Println(c)

	w := woo()
	g := w()
	fmt.Println(g)
	
	sup()
	fmt.Println(sup()())

	tea()
	fmt.Println(tea()())

	shit()
	fmt.Println(shit()())

	wee()
	fmt.Println(wee()())

	space()
	fmt.Println(space()())

	k1 := shell("kelp")
	h1 := k1
	fmt.Println(h1)

	k2 := shell("mail")
	s2 := k2
	fmt.Println(s2) 

	tall()
	fmt.Println(tall()())

	j2 := tall()
	x1 := j2()
	fmt.Println(x1)

	cc := call()
	ff := cc()
	fmt.Println(ff)

	funk()
	fmt.Println(funk()())

	bb := boo()
	bbb := bb()
	fmt.Println(bbb)

	dd := flip("duper")
	ddd := dd
	fmt.Println(ddd)












da := day()
da2 := da()
fmt.Println(da2)


}


func day() func() string {
	return func() string {
		return "something"
	}
}













func cra() func() string {
	return func() string {
		return "smiths"

	}
}

func flip(s string) string {
	return "smeeeeeee"
}

func boo() func () bool {
	return func() bool {
		return false
	}
}

func funk() func() string {
	return func() string {
		return "snowstuff"

	}
}

func call() func() string {
	return func() string {
		return "baseball"
	}
}

func tall() func() int {
	return func() int {
		return 24
	}
}

func shell(t string) string {
	return "shell smell"
}

func space() func() int {
	return func() int {
		return 44
	}
}

func wee() func() string {
	return func() string {
		return "its a small world after all"
	}
}

func shit() func() string {
	return func() string {
		return "uhhhh little kids are frustrating"
	}
}


func tea() func() bool {
	return func() bool {
		return true
	}
}

func sup() func() string {
	return func() string {
		return "tanya programmer"
	}
}

func woo() func() int {
	return func() int {
		return 34
	}
}

func foo() string {
	return "Hello world"
	
}

func bar() func() int {
	return func() int {
		return 42
	}
}

func sam() func() string {
	return func() string {
		return "sup yall"
	}
}

func temp() func() int {
	return func() int {
		return 55
	}
}