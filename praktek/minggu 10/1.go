package main

import "fmt"

const nmax = 1000

type tabArr = [nmax]int

func searching(data tabArr, key int, nSuara int, jumlah *int) {
	i := 0
	for i < nSuara {
		if data[i] == key {
			*jumlah++
		}
		i++
	}
}

func main() {
	// deklarasi variable
	var x, suaraMasuk, i, jumlah int
	var pemilu tabArr
	// input data array
	i = 0
	fmt.Scan(&x)
	for x != 0 {
		suaraMasuk++
		// validasi data
		if x > 0 && x <= 20 {
			pemilu[i] = x
			i++
		}
		fmt.Scan(&x)
	}
	fmt.Printf("Suara masuk: %d\nSuara sah: %d\n", suaraMasuk, i)
	for j := 1; j <= 20; j++ {
		jumlah = 0
		searching(pemilu, j, suaraMasuk, &jumlah)
		if jumlah != 0 {
			fmt.Printf("%d: %d\n", j, jumlah)
		}
	}
}

// 7 19 3 2 78 3 1 -3 18 19 0
