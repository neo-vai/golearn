package main

import "fmt"

func main() {
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Typy of c1: %T\n", c1)
	fmt.Printf("Typy of c2: %T\n", c2)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	fmt.Printf("Typof of c3: %T\n", c3)
	cZero := c3 - c3
	fmt.Println("cZero:", cZero)

	x := 12
	k := 5
	fmt.Println("x:", x)
	fmt.Printf("Type of x: %T\n", x)
	fmt.Println("k:", k)
	fmt.Printf("Type of k: %T\n", k)

	div := 12 / 5
	fmt.Println("div:", div)
	fmt.Printf("Typof of div: %T\n", div)

	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)

	y := 4 / 2.3
	fmt.Println("y:", y)

	divFloat := float64(x) / float64(k)
	fmt.Println("divFloat:", divFloat)
	fmt.Printf("Type of difFloat: %T\n", divFloat)

}
