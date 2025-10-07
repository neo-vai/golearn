package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const MIN = 33
const MAX = 126

func genereateRandomString() string {
	temp := ""
	for i := 0; i < 15; i++ {
		temp += string(rand.Intn(MAX-MIN) + MIN)
	}
	return temp
}

func main() {
	aMap := make(map[string]string)
	aMap["123"] = "123"
	aMap["key"] = "value"

	for i := 0; i < 15; i++ {
		aMap[strconv.Itoa(i)] = genereateRandomString()
	}

	for _, v := range aMap {
		fmt.Println(" # ", v)
	}
}
