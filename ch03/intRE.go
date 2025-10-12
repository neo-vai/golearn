package main

import (
	"fmt"
	"os"
	"regexp"
)

func mutchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[+-]?\d+$`)
	return re.Match(t)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Not enought arrguments")
		return
	}
	fmt.Println(mutchInt(os.Args[1]))
}
