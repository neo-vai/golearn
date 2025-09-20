package main

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	Print([]int{1, 2, 3})
	Print([]float64{1.0, 2.0, 3.0})
	Print([]string{"One", "Two", "Three"})
}
