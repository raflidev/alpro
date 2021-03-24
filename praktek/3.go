package main

import "fmt"

func main() {
	var sks, matkul, jumlah, totalSks int
	var nilai string
	var daftarnilai = map[string]int{"A": 4, "B": 3, "C": 2, "D": 1, "E": 0}
	fmt.Scanln(&matkul)
	for i := 0; i < matkul; i++ {
		fmt.Scanln(&nilai, &sks)
		for sks < 0 || (nilai != "A" && nilai != "B" && nilai != "C" && nilai != "D" && nilai != "E") {
			fmt.Scanln(&nilai, &sks)
		}
		totalSks = totalSks + sks
		jumlah = jumlah + (daftarnilai[nilai] * sks)
	}
	fmt.Println(jumlah / totalSks)

}
