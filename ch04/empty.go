package main

import "fmt"

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1 int
	F2 string
}

func Print(s interface{}) {
	fmt.Println(s)
}

func main() {
	v1 := S1{19, "HEllO"}
	v2 := S2{3238, "BYBY"}
	Print(v1)
	Print(v2)
}
