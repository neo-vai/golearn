package main

import "fmt"

func change(s []string) {
	s[0] = "Changed_by_change-function"
}

func main() {
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)

	var s0 = a[:1]
	fmt.Println(s0)
	s0[0] = "S0"

	var s12 = a[1:3]
	fmt.Println(s12)
	s12[0] = "S12_0"
	s12[1] = "S12_1"

	fmt.Println("a:", a)

	change(s12)
	fmt.Println("a:", a)

	fmt.Println("Capacity of s0:", cap(s0), "| Length of s0:", len(s0))

	s0 = append(s0, "N1")
	s0 = append(s0, "N2")
	s0 = append(s0, "N3")
	a[0] = "-N1"

	s0 = append(s0, "N4")
	fmt.Println("Cap of s0:", cap(s0), "| Len of s0:", len(s0))
	a[0] = "-N1-"
	a[1] = "-N2-"

	fmt.Println("s0:", s0)
	fmt.Println("a:", a)
	fmt.Println("s12", s12)

}
