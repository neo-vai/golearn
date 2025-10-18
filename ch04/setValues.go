package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 int `myapp:"main"`
	F2 string
	F3 float64
}

func main() {
	A := T{1, "Field2", 3.21}

	r := reflect.ValueOf(&A).Elem()

	typrR := r.Type()

	for i := range typrR.NumField() {
		field := typrR.Field(i)
		switch field.Type.Kind() {
		case reflect.String:
			r.Field(i).SetString("NewString")
		case reflect.Int:
			r.Field(i).SetInt(12)
		case reflect.Float64, reflect.Float32:
			r.Field(i).SetFloat(12.1234)
		}

		fmt.Printf("%d: %s : %v\n", i, field.Name, r.Field(i))
		tagValue, ok := field.Tag.Lookup("myapp")
		if ok {
			fmt.Printf("\tTAG: %s\n", tagValue)
		}
	}

}
