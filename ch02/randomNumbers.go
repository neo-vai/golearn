package main

import (
	"fmt"
	"math/rand"
)

func random(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	n := random(0, 100)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Print(random(100, 1000), " ")
	}
}
