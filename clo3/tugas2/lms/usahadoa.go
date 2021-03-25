package main

import "fmt"

var usaha, doa_banyak, total int
var doa_ortu bool
var nilai_algoritma string

func main() {
	fmt.Print("Banyaknya usaha? ")
	fmt.Scanln(&usaha)
	fmt.Print("Banyaknya doa? ")
	fmt.Scanln(&doa_banyak)
	fmt.Print("Doa orangtua? ")
	fmt.Scanln(&doa_ortu)
	fmt.Print("Nilai Algoritma? ")
	fmt.Scanln(&nilai_algoritma)

	BacaData(&usaha, doa_banyak, doa_ortu, nilai_algoritma)
}

func BacaData(usaha *int, jumlahDoa int, doaOrtu bool, nilai string) {
	TabungDoaUsaha(*usaha, jumlahDoa, &total)
	total = TabungUsahaDoa(doaOrtu, &total)
	total = HasilNilaiAlpro(nilai, &total)
	fmt.Println(HasilAkhir(total))
}

func TabungDoaUsaha(usaha int, doa int, total *int) {
	*total = doa + usaha
}
func TabungUsahaDoa(doa bool, total *int) int {
	if doa == true {
		return *total * 2
	} else {
		return *total
	}
}
func HasilNilaiAlpro(nilai string, total *int) int {
	if nilai == "A" {
		*total = *total - 150
	} else if nilai == "B" {
		*total = *total - 130
	} else if nilai == "C" {
		*total = *total - 100
	} else if nilai == "D" {
		*total = *total - 0
	} else if nilai == "E" {
		*total = *total - 0
	}
	return *total
}
func HasilAkhir(poin int) string {
	if poin >= 130 {
		return "Lulus langsung dapat kerja gaji 2 digit"
	} else if poin >= 50 && poin < 130 {
		return "Lulus langsung dapat kerja"
	} else {
		return "Jangan lelah berdoa dan berusaha, tidak ada yang sia – sia dari usaha dan doa"
	}
}
