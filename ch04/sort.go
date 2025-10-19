package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

type StringHelper struct {
	F1 string
	F2 string
	F3 string 
}

type SSHelper []StringHelper

func (SSH SSHelper) Len() int {
	return len(SSH)
}

func (SSH SSHelper) Less(i, j int) bool {
	slen := utf8.RuneCountInString
	return slen(SSH[i].F1) + slen(SSH[i].F2) + slen(SSH[i].F3) < slen(SSH[j].F1) + slen(SSH[j].F2) + slen(SSH[j].F3)
}

func (SSH SSHelper) Swap(i, j int) {
	SSH[i], SSH[j] = SSH[j], SSH[i]
}

func main() {
	SSH := SSHelper	{
		StringHelper{
			F1: "Hello",
			F2: "H",
			F3: "YAJSKDL;FKASJDFLKJALSDFASKLDFJ",
		},
		StringHelper{
			F1: "Hello My Friend",
			F2: "H",
			F3: "Y",
		},
	}
	fmt.Println("Before:", SSH)
	sort.Sort(SSH)
	fmt.Println("After", SSH)
	sort.Sort(sort.Reverse(SSH))
	fmt.Println("Reverse", SSH)
}
