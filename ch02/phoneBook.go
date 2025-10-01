package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}
var MIN = 0
var MAX = 26

func Search(key string) *Entry {
	for i, v := range data {
		if v.Tel == key {
			return &data[i]
		}
	}
	return nil
}

func List() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func GetString(l int64) string {
	startChar := "A"
	tmp := ""
	var i int64 = 1
	for {
		myRand := Random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		tmp = tmp + newChar
		if i == l {
			break
		}
		i++
	}
	return tmp
}

func populate(n int, s []Entry) {
	for i := 0; i < n; i++ {
		name := GetString(4)
		surname := GetString(5)
		n := strconv.Itoa(Random(800000000, 899999999))
		data = append(data, Entry{name, surname, n})
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	n := 100
	populate(n, data)
	fmt.Printf("Data has %d entries.\n", len(data))

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := Search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		List()
	default:
		fmt.Println("Not a valid operation")
	}
}
