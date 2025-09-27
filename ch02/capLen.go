package main

import "fmt"

func main() {
	a := make([]int, 4)
	// Len == Cap
	fmt.Printf("Len: %d, Cap: %d\n", len(a), cap(a))
	b := []int{0, 1, 2, 3, 4}
	// Len == Cap
	fmt.Printf("Len: %d, Cap: %d\n", len(b), cap(b))

	// Len == Cap
	// 4 == 4
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)
	// Len == 4+1 == 5 -> Cap != 4; Cap *= 2 == 8
	aSlice = append(aSlice, 5)
	fmt.Println(aSlice)
	fmt.Printf("Len: %d, Cap: %d\n", len(aSlice), cap(aSlice))

	// Len == 5 + 4 == 9 -> Cap != 8; Cap *= 2 == 16
	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)

	fmt.Println(aSlice)
	fmt.Printf("Len: %d, Cap: %d\n", len(aSlice), cap(aSlice))

}
