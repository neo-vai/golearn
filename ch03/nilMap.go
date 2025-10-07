package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic: ", r)
		}
	}()

	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("aMap:", aMap)
	aMap = nil

	fmt.Println("aMap:", aMap)
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}

	aMap["test"] = 1

	aMap = nil
	aMap["test"] = 1
	fmt.Println("aMap", aMap)
}
