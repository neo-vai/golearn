package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("Both arguments are zero")
	}
	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("Both arguments are zero: a %d and b %d. UserdID: %d", a, b, os.Getuid())
	}
	return nil
}

func main() {

	// First test

	if err := check(0, 1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Error ended normaly!")
	}

	if err := check(0, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Error ended normaly!")
	}

	// Second test

	if err := formattedError(-123, 112); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Error ended normaly!")
	}

	if err := formattedError(0, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Error ended normaly!")
	}

	// Three test

	if i, err := strconv.Atoi("-123"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Recognized umber:", i)
	}

	if i, err := strconv.Atoi("-F123"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Recognized umber:", i)
	}

}
