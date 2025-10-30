package main

import "fmt"

type IntA interface {
	foo()
}

type IntB interface {
	boo()
}

type IntC interface {
	IntA
	IntB
}

func processA(s IntA) {
	fmt.Printf("%T\n", s)
}

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
}

type c struct {
	A a
	B b
}

type compose struct {
	field1 int
	a
}

func (varC c) foo() {
	fmt.Println("Foo precessing...", varC)
}

func (varC c) boo() {
	fmt.Println("Boo processing...", varC)
}
