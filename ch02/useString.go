package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	f("EqualFold: %v\n", s.EqualFold("NeoVai", "neovai"))
	f("EqualFold: %v\n", s.EqualFold("NaVai", "NeoVais"))

	f("Index: %v\n", s.Index("Michail", "ha"))
	f("Index: %v\n", s.Index("Nichail", "Ha"))

	f("HasPrefix: %v\n", s.HasPrefix("Michalis", "Mi"))
	f("HasPrefix: %v\n", s.HasPrefix("Michalis", "mi"))
	f("HasSiffix: %v\n", s.HasSuffix("Michalis", "is"))
	f("HasSiffix: %v\n", s.HasSuffix("Michalis", "IS"))

	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	fmt.Println(t)
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))
	fmt.Println(t)

	f("%s\n", s.Split("abcd efg", ""))
	f("%s\n", s.Split("abcd efg", " "))
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", " ", "", 2))

	f("SplitAfter: %s\n", s.SplitAfter("a,b,c", ","))

	f("Trim: %s\n", s.Trim("...123123123...", "."))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	f("TrimFunc: %s\n", s.TrimFunc("123asdf123asd", trimFunction))

}
