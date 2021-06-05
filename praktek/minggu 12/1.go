package main

import "fmt"

const nmax = 100

type Pemain struct {
	namaDepan, namaBelakang string
	goal, assist            int
}

type tabPemain = [nmax]Pemain

func inputPemain(baris *int, dataPemain *tabPemain) {
	fmt.Scanln(*&baris)
	for i := 0; i < *baris; i++ {
		fmt.Scanln(&dataPemain[i].namaDepan, &dataPemain[i].namaBelakang, &dataPemain[i].goal, &dataPemain[i].assist)
	}
}

func sorting(baris int, P *tabPemain) {
	var pass, idx, i int
	var temp Pemain
	pass = 1
	for pass <= baris-1 {
		idx = pass - 1
		i = pass
		for i < baris {
			if P[idx].goal < P[i].goal {
				idx = i
			} else if P[idx].goal == P[i].goal {
				if P[idx].assist < P[i].assist {
					idx = i
				}
			}
			i++
		}

		temp = P[pass-1]
		P[pass-1] = P[idx]
		P[idx] = temp
		pass = pass + 1
	}
}

func cetakPemain(baris int, P tabPemain) {
	for i := 0; i < baris; i++ {
		fmt.Println(P[i].namaDepan, P[i].namaBelakang, P[i].goal, P[i].assist)
	}
}

func main() {
	var dataPemain tabPemain
	var baris int
	inputPemain(&baris, &dataPemain)
	sorting(baris, &dataPemain)
	fmt.Print("\n")
	cetakPemain(baris, dataPemain)
}

// 9
// Bruno Fernandes 7 3
// Jamie Vardy 8 1
// Dominic Calvert-Lewin 10 2
// Son Heung-Min 9 2
// Harry Kane 7 2
// Ollie Watkins 6 1
// Patrick Bamford 7 1
// Callum Wilson 7 0
// Mohamed Salah 8 2
