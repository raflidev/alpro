package main

import "fmt"

func main() {
	var i, j, digit, angka, hasil int
	fmt.Scanln(&digit)
	i = 0
	for i < digit {
		fmt.Scanln(&angka)
		if angka > -1 && angka < 10 {
			j = 0
			for j < angka {
				angka = angka * 10
				j++
			}
			hasil = hasil + angka
			i++
		}
	}
	fmt.Println(hasil)
}
