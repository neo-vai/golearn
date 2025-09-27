package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func deleteV1[T any](slice []T, i int) ([]T, error) {
	if i < 0 || i >= len(slice) {
		return slice, fmt.Errorf("index %d out of range [0,%d)", i, len(slice))
	}
	return append(slice[:i], slice[i+1:]...), nil
}

func deleteV2[T any](slice []T, i int) ([]T, error) {
	if i < 0 || i >= len(slice) {
		return slice, fmt.Errorf("index %d out of range [0,%d)", i, len(slice))
	}
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1], nil
}

func main() {

	lenArgs := len(os.Args)

	// Valid
	if lenArgs < 4 || (os.Args[1] != "v1" && os.Args[1] != "v2") || !func() bool {
		if e, err := strconv.Atoi(os.Args[lenArgs-1]); err == nil && e < lenArgs-3 {
			return true
		}
		return false
	}() {
		exeName := ""
		exePath, err := os.Executable()
		if err == nil {
			exeName = filepath.Base(exePath)
		}

		fmt.Printf("Usage: %s v1|v2 arg1 arg2 arg3... [delete-element-index]\n", exeName)
		return
	}

	// v1 or v2 delte method
	var deleteFunc func([]string, int) ([]string, error)

	// slice of elements
	slice := os.Args[2 : lenArgs-1]

	// delete element index
	i, err := strconv.Atoi(os.Args[lenArgs-1])
	if err != nil {
		fmt.Println("Delete element indext is not int")
		return
	}

	switch os.Args[1] {
	case "v1":
		deleteFunc = deleteV1
	case "v2":
		deleteFunc = deleteV2
	default:
		panic("Incorrect delete func name")
	}

	slice, err = deleteFunc(slice, i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(slice)
}
