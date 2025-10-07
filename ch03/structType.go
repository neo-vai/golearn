package main

import "fmt"

type User1 struct {
	Name string
	Age  int
}

type User2 struct {
	Name string
	Age  int
}

type User3 struct { // You have to chage field order
	Age  int
	Name string
}

// Field order ^

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Compile error, another types: ", r)
		}
	}() // doesnt work, cuz it`s not panic, It`s compile error

	user1 := User1{"Bob", 13}
	user2 := User2(user1)
	user3 := User3(user2) // error, Not panic, Comiple error
	// struct {Name string, Age int} == struct {Name string, Age int} != struct {Age int, Name string}
	fmt.Println(user1, user2, user3)
}
