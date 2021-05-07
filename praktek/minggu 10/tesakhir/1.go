package main

import "fmt"

// deklarasi range array
const nmax = 1000

// membuat array
type tabPemilu = [nmax]int

func main() {
	// deklarasi variable
	var dataPemilu tabPemilu
	var x, i, suaraMasuk, jumlah int
	i = 0
	// input voting
	fmt.Scan(&x)
	for x != 0 {
		suaraMasuk++
		if x > 0 && x <= 20 {
			dataPemilu[i] = x
			i++
		}
		fmt.Scan(&x)
	}
	// output suara masuk dan suara sah
	fmt.Printf("Suara masuk: %d\nSuara sah: %d\n", suaraMasuk, i)
	z := 0
	for j := 1; j <= 20; j++ {
		jumlah = 0
		// calon kandidat
		var pilihan = [2]string{"Ketua RT", "Wakil RT"}
		searching(dataPemilu, j, suaraMasuk, &jumlah)
		// menampilkan hasil lebih dari 1
		if jumlah > 1 {
			fmt.Printf("%s: %d\n", pilihan[z], j)
			z++
		}
	}
}

func searching(data tabPemilu, key int, jumlahSuara int, jumlah *int) {
	// mencari kesamaan data
	i := 0
	for i < jumlahSuara {
		if data[i] == key {
			*jumlah++
		}
		i++
	}
}

// 7 19 3 2 78 3 1 -3 18 19 0
