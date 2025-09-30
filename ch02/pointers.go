package main

import "fmt"

type aStructure struct {
	fild1 complex64
	fild2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointer(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 10.5
	fmt.Println("Memory addres of f:", &f)

	fP := &f
	fmt.Println("Memory addres of f:", fP)
	fmt.Println("Vaule of f:", *fP)

	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)

	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)

	xx := bothPointer(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)

	var k *aStructure
	fmt.Println(k)

	if k == nil {
		k = new(aStructure)
	}

	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil!")
	}
}
