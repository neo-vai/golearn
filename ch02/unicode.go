package main

import (
	"fmt"
	"os"
	"unicode"
)

func UnicodePrint(sL string) {
	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c", sL[i])
		} else {
			fmt.Print("\033[31mX\033[0m")
		}
	}
	println()
}

func main() {
	if len(os.Args) == 1 {
		UnicodePrint("He\033llo wo\nld")
	} else {

		for _, v := range os.Args[1:] {
			UnicodePrint(v)
		}
	}
}
