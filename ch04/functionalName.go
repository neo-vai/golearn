package main

import "fmt"

type Something interface {
	Touch()
}

type File struct {
}

func (f *File) Touch() {
	fmt.Println("AAAAAA")
}

func TryToTouch(s Something) {
	fmt.Println()
	for range 12 {
		s.Touch()
	}
	fmt.Println()
}

func main() {
	A := File{}
	TryToTouch(&A)
	A.Touch()
}
