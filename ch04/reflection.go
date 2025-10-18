package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"My Value", -12.22, Secret{"neovai", "JDK"}}

	reflectObject(A)

}

func reflectObject(i any) {
	r := reflect.ValueOf(i)

	iType := r.Type()

	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", r.NumField(), r.Type())

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s", iType.Field(i).Name)
		fmt.Printf("\twith type: %s", r.Field(i).Type())
		fmt.Printf("\tand value _%v_", r.Field(i).Interface())
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Struct {
			fmt.Print(" : ", r.Field(i).Type())
		}

		fmt.Println()

	}

}
