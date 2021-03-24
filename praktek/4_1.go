package main

import "fmt"

func main() {
	var AR, AT, ronde int
	var penyaji, penebak, pemenang string
	penyaji = "A"
	penebak = "B"
	ronde = 0
	for true {
		ronde = ronde + 1
		fmt.Println("Ronde ", ronde, ":")
		fmt.Print(penyaji, " - Masukkan angka rahasia : ")
		fmt.Scanln(&AR)
		if AR == -101 {
			fmt.Println("Permainan Selesai")
			break
		}
		for i := 1; i <= 3; i++ {
			fmt.Print(penebak, " - Masukkan angka tebakan ke-", i, " : ")
			fmt.Scanln(&AT)
			if AT == AR {
				pemenang = penebak
				break
			} else {
				pemenang = penyaji
			}
		}
		fmt.Println(pemenang, "Adalah pemenangnya")
		if pemenang == "A" {
			penebak = "B"
			penyaji = "A"
		} else {
			penebak = "A"
			penyaji = "B"
		}
	}

}
