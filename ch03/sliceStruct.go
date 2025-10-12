package main

import (
	"fmt"
	"strconv"
)

type record struct {
	Filed1 int
	Filed2 string
}

func main() {
	S := []record{}
	for i := 0; i < 10; i++ {
		text := "text" + strconv.Itoa(i)
		temp := record{Filed1: i, Filed2: text}
		S = append(S, temp)
	}

	fmt.Println("Index 2:", S[2].Filed1, S[2].Filed2)
	fmt.Println("Number of sctructures:", len(S))
	sum := 0
	for _, k := range S {
		sum += k.Filed1
	}
	fmt.Println("Sum:", sum)
}
