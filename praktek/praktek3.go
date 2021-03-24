package main

import "fmt"

func main() {
	var r, luasLingkaran float64
	fmt.Scanln(&r)
	luasLingkaran = 22.0 / 7.0 * r
	fmt.Println("Luas lingkaran dengan jari-jari =", r, "adalah", luasLingkaran)
}
