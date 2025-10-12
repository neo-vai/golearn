package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Hot enough arguments")
		return
	}
	fmt.Println(matchNameSur(os.Args[1]))
}
