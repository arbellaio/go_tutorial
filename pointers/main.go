package main

import "fmt"

type Foo struct {
	bar int
}

func main() {

	var foo *Foo
	fmt.Println(foo)

	foo = new(Foo)
	fmt.Println(foo)
	fmt.Println(*foo) //to get value with struct
	fmt.Println((*foo).bar) //to get value only
	(*foo).bar = 45
	fmt.Println((*foo).bar) //to get value only

	foo.bar = 47
	fmt.Println(foo.bar)

	var a int = 12
	var b *int = &a //here *represents pointer and & represents address
	// of memory location
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) //dereferencing the memory location meaning
	//getting value from location
}
