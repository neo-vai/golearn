package main

import "fmt"

type ar2x2 [2][2]int

func Add(a, b ar2x2) ar2x2 {
	c := ar2x2{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func (a *ar2x2) Add(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] += b[i][j]
		}
	}
}

func (a *ar2x2) Substract(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] -= b[i][j]
		}
	}
}

func (a *ar2x2) Multiply(b ar2x2) {
	result := ar2x2{}

	result[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	result[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	result[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	result[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]
	*a = result
}

func main() {
	a := ar2x2{{0, 1}, {2, 3}}
	b := ar2x2{{3, 2}, {1, 0}}
	a.Add(b)
	fmt.Println(a)
	b.Substract(b)
	fmt.Println(b)
	a = ar2x2{{0, 1}, {2, 3}}
	b = ar2x2{{3, 2}, {1, 0}}

	a.Multiply(b)
	fmt.Println(a)
}
