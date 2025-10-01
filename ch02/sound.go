package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var SYMBOL time.Duration
var SPACE time.Duration

func runLine(line string) error {
	isD := false
	tempD := ""
	for _, v := range line {
		switch v {
		case ' ':
			fmt.Print(" ")
			time.Sleep(SPACE)
		case '%':
			if len(tempD) > 0 && isD {
				isD = false

				number, err := strconv.Atoi(tempD[1:])
				if err != nil {
					return errors.New("Time sleep is not number: " + tempD[1:])
				}

				switch tempD[0] {
				case 'd':
					time.Sleep(time.Duration(number) * time.Millisecond)
				case 's':
					SYMBOL = time.Duration(number) * time.Millisecond
				default:
					return errors.New("Unknows symbol: " + string(tempD[0]))
				}

				tempD = ""
			} else {
				isD = true
			}
		default:
			if isD {
				tempD += string(v)
				continue
			}
			time.Sleep(SYMBOL)
			fmt.Printf(string(v))
		}
	}
	fmt.Println()
	return nil
}

func RunSound(words []string) {
	for _, v := range words {
		runLine(v)
	}
}

func main() {
	SPACE = time.Millisecond * 100
	SYMBOL = time.Millisecond * 80

	b := []string{
		"\nYou are an %s0%\033[31m%s" + strconv.Itoa(int(SYMBOL.Milliseconds())) + "%IDIOT%s0%\033[0m%s" + strconv.Itoa(int(SYMBOL.Milliseconds())) + "%%d200%",
		"Hahahahahaha!%d500%",
		"Hahahaha!%d2200%",
	}

	time.Sleep(3500 * time.Millisecond)

	RunSound(b)
	RunSound(b)
	time.Sleep(200 * time.Millisecond)
	RunSound(b)
	RunSound(b)
	time.Sleep(100 * time.Millisecond)
}
