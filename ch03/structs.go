package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Year    int
}

func zeroS() Entry {
	return Entry{}
}

func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		Y = 2000
	}
	return Entry{Name: N, Surname: S, Year: Y}
}

func zeroPtoS() *Entry {
	return &Entry{}
}

func initPtoS(N, S string, Y int) *Entry {
	a := initS(N, S, Y)
	return &a
}

func main() {
	s1 := zeroS()
	p1 := zeroPtoS()
	fmt.Println("s1:", s1, "| p1:", p1)
			
	s2 := initS("Neo", "Vai", 2000)
	p2 := initPtoS("Neo", "Vai", 2000)

	fmt.Println("s2:", s2, "| p2:", *p2)

	fmt.Println("Year:", s1.Year, p1.Year, s2.Year, p2.Year)

	pS := new(Entry)
	fmt.Println("pS:", pS.Name, "\n", pS.Surname, "\n", pS.Year)
}
