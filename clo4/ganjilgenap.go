package main

import "fmt"

func main() {
	var angka int
	var genap, ganjil [1]int
	fmt.Scanln(&angka)
	if angka%2 == 0 {
		genap[0] = angka
	} else {
		ganjil[0] = angka
	}
	fmt.Print("genap: ")
	fmt.Println(genap)
	fmt.Print("ganjil: ")
	fmt.Println(ganjil)
}
