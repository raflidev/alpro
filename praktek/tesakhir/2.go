package main

import "fmt"

func main() {
	var i, bil, jumlah, banyak int
	i = 1
	bil = 1
	for bil != 0 {
		fmt.Print("Bilangan ke-", i, ": ")
		fmt.Scanln(&bil)
		for bil < 0 {
			fmt.Print("Bilangan ke-", i, ": ")
			fmt.Scanln(&bil)
		}
		if bil%i == 0 && bil != 0 {
			jumlah = jumlah + bil
			banyak += 1
		}
		i += 1
	}
	fmt.Println("Banyaknya bilangan:", banyak)
	fmt.Println("Total penjumlahan:", jumlah)
}
