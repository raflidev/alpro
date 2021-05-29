package main

import "fmt"

const NMAX = 200

type angka = [NMAX]int

func ganjilGenap(angka int) string {
	if angka%2 == 0 {
		return "genap"
	}
	return "ganjil"
}

func terbesar(tabAngka *angka, n int, hasil *int) {
	max := 0
	idx := 0
	for i := 0; i < n; i++ {
		if tabAngka[i] > max {
			max = tabAngka[i]
			idx = i
		}
	}
	*hasil = tabAngka[idx]
	tabAngka[idx] = 0
}

func terkecil(tabAngka *angka, n int, hasil *int) {
	min := tabAngka[0]
	idx := 0
	for i := 0; i < n; i++ {
		if min > tabAngka[i] {
			min = tabAngka[i]
			idx = i
		}
	}
	*hasil = tabAngka[idx]
	tabAngka[idx] = 9999
}

func main() {
	var a, b, max1, max2, min1, min2 int
	var tabAngka angka
	fmt.Scanln(&a, &b)

	for i := 0; i < a; i++ {
		fmt.Scan(&tabAngka[i])
	}
	if ganjilGenap(b) == "ganjil" {
		terbesar(&tabAngka, a, &max1)
		terbesar(&tabAngka, a, &max2)
		fmt.Println(max1, max2)
	} else {
		terkecil(&tabAngka, a, &min1)
		terkecil(&tabAngka, a, &min2)
		fmt.Println(min1, min2)
	}
}

// 18 3
// 12 3 45 67 8 90 1 23 42 56 7 89 30 126 14 5 679 90

// 18 100
// 12 3 45 67 8 90 679 23 42 56 7 89 30 126 14 5 2 90
