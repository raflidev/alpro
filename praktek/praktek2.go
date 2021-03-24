package main

import "fmt"

func main() {
	var s string
	var a, b int
	fmt.Scanln(&s, &a, &b)
	var hasilPenjumlahan = a + b
	fmt.Println("Kata =", s)
	fmt.Println("Jumlah =", hasilPenjumlahan)
}
