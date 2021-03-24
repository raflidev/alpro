package main

import "fmt"

func main() {
<<<<<<< HEAD
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

=======
	var winner, player, temp string
	var i, ronde int
	var nilai, answer int

	winner = "A"
	player = "B"
	ronde = 1
	fmt.Println("Ronde", ronde, ":")
	fmt.Print(winner, " - Masukan angka rahasia: ")
	fmt.Scan(&nilai)
	for nilai != -101 {
		i = 1
		fmt.Print(player, "- Masukan angka tebakan ke-", i, ":")
		fmt.Scan(&answer)
		for i < 3 && answer != nilai {
			i = i + 1
			fmt.Print(player, "- Masukan angka tebakan ke-", i, ":")
			fmt.Scan(&answer)
		}
		if nilai == answer {
			temp = winner
			winner = player
			player = temp
		}
		fmt.Println(winner, "pemenangnya")
		ronde = ronde + 1
		fmt.Println("Ronde", ronde, ":")
		fmt.Print(winner, " - Masukan angka rahasia: ")
		fmt.Scan(&nilai)
	}
	fmt.Println("permainan selesai")
>>>>>>> a186180cdf3afbf0b4d511393b6ccb906ca9b9be
}
