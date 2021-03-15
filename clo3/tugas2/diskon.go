package main

import "fmt"

func main() {
	var belanja, diskon int
	var member bool

	fmt.Scanln(&belanja, &member)
	diskon = ambilDiskon(belanja, member)
	fmt.Println("Total Bayar :", belanja-((belanja*diskon)/100))
}

func ambilDiskon(belanja int, member bool) int {
	var diskon int
	if belanja > 100000 {
		if member {
			diskon = 10
		} else {
			diskon = 5
		}
	}
	return diskon
}
