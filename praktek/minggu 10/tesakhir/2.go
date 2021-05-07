package main

import "fmt"

// deklarasi range array
const nmax = 1000000

// deklarasi array
var data [nmax]int

func main() {
	// deklarasi variable
	var n, k int
	// menginput angka ke variable n dan k
	fmt.Scanln(&n, &k)
	// memanggil procedure isiArray dengan parameter n
	isiArray(n)
	// output posisi
	// percabangan, jika posisi tidak ditemukan
	posisi := posisi(n, k)
	if posisi == -1 {
		fmt.Println("Tidak ada")
	} else {
		// posisi ditemukan
		fmt.Println(posisi)
	}
}

func isiArray(n int) {
	// menginput data angka ke array data
	i := 0
	for i < n {
		fmt.Scan(&data[i])
		i++
	}
}

func posisi(n, k int) int {
	// mencari posisi index didalam array data
	var index int
	i := 0
	found := false
	for i < n && !found {
		if data[i] == k {
			index = i
			found = true
			return index
		}
		i++
	}
	return -1
}

// 12 534
// 1 3 8 16 32 123 323 323 534 543 823 999
