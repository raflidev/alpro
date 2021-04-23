package main

import "fmt"

func main() {
	var nama1, nama2 string
	var soal1, soal2, skor1, skor2 int
	fmt.Scan(&nama1)
	hitungSkor(&soal1, &skor1)
	fmt.Scan(&nama2)
	hitungSkor(&soal2, &skor2)

	if soal1 == soal2 {
		if skor1 < skor2 {
			fmt.Println(nama1, soal1, skor1)
		} else {
			fmt.Println(nama2, soal2, skor2)
		}
	} else if soal1 > soal2 {
		fmt.Println(nama1, soal1, skor1)
	} else {
		fmt.Println(nama2, soal2, skor2)
	}

}

func hitungSkor(soal, skor *int) {
	var waktu int
	i := 0
	for i < 8 {
		fmt.Scan(&waktu)
		if waktu != 301 {
			*soal += 1
			*skor += waktu
		}
		i++
	}
}
