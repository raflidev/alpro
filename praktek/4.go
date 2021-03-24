package main

import "fmt"

func main() {
	var rahasia, tebakan, salahtebak, ronde int
	var player = [2]string{"A", "B"}
	var playerNow, enemy string
	playerNow = player[0]
	enemy = player[1]
	for true {
		ronde = ronde + 1
		salahtebak = 0

		fmt.Print("Ronde ", ronde, ":\n")
		fmt.Print(playerNow, " - masukan angka rahasia: ")
		fmt.Scan(&rahasia)
		if rahasia >= 1 && rahasia <= 20 {
			for i := 1; i <= 3; i++ {
				fmt.Print(enemy, " - masukan angka tebakan ke-", i, ": ")
				fmt.Scan(&tebakan)
				if tebakan != rahasia {
					salahtebak = salahtebak + 1
				} else {
					fmt.Println(enemy, "adalah pemenangnya \n\n")
					if enemy == "A" {
						playerNow = player[0]
						enemy = player[1]
					} else {
						playerNow = player[1]
						enemy = player[0]

					}
					break
				}

			}

			if salahtebak == 3 {
				fmt.Println(playerNow, "adalah pemenangnya \n\n")
				if playerNow != "A" {
					playerNow = player[1]
					enemy = player[0]
				} else {
					playerNow = player[0]
					enemy = player[1]
				}

			}

		} else if rahasia == -101 {
			fmt.Println("Permainan selesai")
			break
		}
	}
}
