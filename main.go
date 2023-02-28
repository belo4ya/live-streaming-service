package main

import (
	"fmt"
)

type Foo struct {
	A   string
	B   int
	Foo *Foo
}

type Bar struct {
	A   string
	B   int
	Foo *Foo
}

func main() {
	foo := Foo{A: "string", B: 10, Foo: &Foo{A: "nested-string", B: 20}}
	fmt.Printf("%+v\n", foo)
	fmt.Printf("%+v\n", Bar(foo))

	fooPointer := &foo
	barPointer := Bar(*fooPointer)
	fmt.Printf("\n%+v\n", fooPointer)
	fmt.Printf("%+v\n", &barPointer)
}
