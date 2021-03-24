package main

import "fmt"

func main() {
	var n, jumlah int
	fmt.Scanln(&n)
	jumlah = prima(n, jumlah)
	fmt.Println("Banyak prima:", jumlah)
}

func prima(n, jumlah int) int {
	for n > 0 {
		for i := 2; i <= n; i++ {
			if n%i == 0 {
				jumlah = jumlah + 1
			}
		}
		fmt.Scanln(&n)
	}
	return jumlah
}
