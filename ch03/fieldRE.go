package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[+]?\d+$`)
	return re.Match(t)
}

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func mathRecord(s string) bool {
	fields := strings.Split(s, ",")
	if len(fields) != 3 {
		return false
	} else if !matchNameSur(fields[0]) {
		return false
	} else if !matchNameSur(fields[1]) {
		return false
	}

	return matchTel(fields[2])

}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Not enough arguments")
		return
	}

	fmt.Println(mathRecord(os.Args[1]))
}
