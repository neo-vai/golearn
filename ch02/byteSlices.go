package main

import "fmt"

func main() {
	b := make([]byte, 12)
	fmt.Println("Byte slices:", b)

	b = []byte("Hello My firend €")
	fmt.Println("Byte lices:", b)

	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))

	fmt.Println("Length of b:", len(b))
}
