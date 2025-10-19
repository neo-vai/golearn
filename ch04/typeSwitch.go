package main

import (
	"fmt"
	"reflect"
)

type SecretStruct struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 SecretStruct
}

func TestStruct(x interface{}) {
	switch T := x.(type) {
	case SecretStruct:
		fmt.Println("Secret Type")
	case Entry:
		fmt.Println("Enty Type")
	default:
		fmt.Printf("Not supportet type: %T\n", T)
	}
}

func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}

func ReflectTypeTest(x any) {
	v := reflect.ValueOf(x)
	vType := v.Type()
	switch vType.Kind() {
	case reflect.Struct:
		fmt.Println("Its a structure")
	case reflect.String:
		fmt.Println("Its a string")
	case reflect.Int:
		fmt.Println("Its an integer")
	case reflect.Float64:
		fmt.Println("Its a FUCKING float")
	case reflect.Func:
		fmt.Println("Are you stupid?! Its a FUCKING func")
	default:
		fmt.Println("Unknown value")
	}
}

func main() {
	ent := Entry{
		F1: 3601,
		F2: "My Sting***",
		F3: SecretStruct{
			"MYKEY63238",
		},
	}
	ReflectTypeTest(ent)
	ReflectTypeTest(12)
	ReflectTypeTest("HELLO MY FUCKING FRIEND")
	ReflectTypeTest(12.12)
	ReflectTypeTest(func(){})
	ReflectTypeTest(float32(12.2))

	TestStruct(ent)
	TestStruct(ent.F3)
	TestStruct(ent.F1)

	Learn(12)
	Learn("Hello my teacher")
	Learn(func(a, b, c, d int)int{return a + b + c + d})
}
