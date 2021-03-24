package main

import "fmt"

func main() {
	// var n, nilai int
	var n, nilai, jumlah, rata2 float64
	fmt.Scanln(&nilai)
	for nilai != -1 {
		n = n + 1
		jumlah = jumlah + nilai
		fmt.Scanln(&nilai)
	}

	if n == 0 {
		rata2 = 0.0
	} else {
		rata2 = jumlah / n
	}

	fmt.Println(rata2)
}
