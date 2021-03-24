package main

import "fmt"

func main() {
	var i, nilaiMax, nilai, abc, n int
	fmt.Scanln(&n)
	i = 0
	nilaiMax = 0
	for i < n {
		fmt.Scan(&nilai)
		if nilai > nilaiMax {
			abc = 1
			nilaiMax = nilai
		} else if nilai == nilaiMax {
			abc++
		}
		i++
	}

	fmt.Println("Nilai terbesar adalah", nilaiMax, "yang diperoleh", abc, "orang mahasiswa")
}
