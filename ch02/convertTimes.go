package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need input __02 01 2006 15:04 MST__ format")
		return
	}

	input := os.Args[1]
	now, err := time.Parse("02 01 2006 15:04 MST", input)
	if err != nil {
		fmt.Println(err)
		return
	}

	loc, _ := time.LoadLocation("Local")
	fmt.Printf("Current Location: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Europe/London")
	fmt.Printf("London time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Tokyo")
	fmt.Printf("Tokyo time: %s\n", now.In(loc))

}
