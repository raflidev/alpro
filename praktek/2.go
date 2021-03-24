package main

import "fmt"

func main() {
	var i, n, bobot, sks, jumlah, jumlahsks int
	var nilai string

	fmt.Scanln(&n)
	i = 1
	for i <= n {
		fmt.Scanln(&nilai, &sks)
		for (sks < 0 || nilai != "A" && nilai != "B" && nilai != "C" && nilai != "D" && nilai != "E") {
			fmt.Scanln(&nilai, &sks)
		}
		if nilai == "A" {
			bobot = 4
		} else if nilai == "B" {
			bobot = 3
		} else if nilai == "C" {
			bobot = 2
		} else if nilai == "D" {
			bobot = 1
		} else {
			bobot = 0
		}
		jumlah = jumlah + (bobot * sks)
		jumlahsks += sks
		i++
	}
	fmt.Println(jumlah / jumlahsks)

}
