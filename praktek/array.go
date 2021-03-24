package main

import "fmt"

func main() {
	var buah = [...]string{"Apel", "Mangga", "Jeruk", "Melon"}

	// var buah = make([]string, 2)
	// buah[0] = "Apel"
	// buah[1] = "Mangga"

	for _, buahan := range buah {
		fmt.Printf("nama buah: %s\n", buahan)
	}
}
