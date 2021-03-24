package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var seed, saya int64
	fmt.Scanf("Angka rahasia: %d", &seed)
	fmt.Scanf("Anda: %d", &saya)
	fmt.Printf("Danang: %d", rand.Seed(seed))

	fmt.Println("Satu nilai 1...6:", rand.Intn(6)+1)
}
