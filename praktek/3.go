package main

import "fmt"

func main() {
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
}
