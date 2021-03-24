package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var seed int64
	var score, saya, dadang int
	fmt.Print("Angka Rahasia: ")

	fmt.Scanln(&seed)
	fmt.Print("Anda: ")
	fmt.Scanln(&saya)
	rand.Seed(seed)
	dadang = rand.Intn(6) + 1
	fmt.Println("Dadang: ", dadang)
	score = rand.Intn(6) + 1
	if saya == dadang {
		fmt.Println("Nilai dadu", score, "Seri")
	} else if score == saya {
		fmt.Println("Nilai dadu", score, "Pemenang adalah Anda")
	} else if score == dadang {
		fmt.Println("Nilai dadu", score, "Pemenang adalah Dadang")
	} else {
		fmt.Println("Nilai dadu", score, "Tidak ada Pemenang")
	}
}
