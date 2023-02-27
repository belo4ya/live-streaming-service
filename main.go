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
	fmt.Printf("%+v", foo)
	fmt.Printf("%+v", Bar(foo))
}
