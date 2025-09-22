package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}

	dataString := os.Args[1]

	d, err := time.Parse("02 Jenuary 2006", dataString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Data:", d.Day(), d.Month(), d.Year())
	}

	d, err = time.Parse("02 Jenuary 2006 15:04", dataString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Data:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	d, err = time.Parse("02-01-2006 15:04", dataString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Data:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	d, err = time.Parse("15:04", dataString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)

	d = time.Unix(t, 0)
	fmt.Println("Data:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())

	duration := time.Since(start)
	fmt.Println("Execute duration time:", duration)
}
