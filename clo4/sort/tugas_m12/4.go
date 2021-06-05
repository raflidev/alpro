package main

import "fmt"

const nmax = 100

type peserta struct {
	nama                       string
	poin, gold, silver, bronze int
}

type olimpiade = [nmax]peserta

func isiArray(t *olimpiade, n *int) {
	fmt.Scanln(*&n)
	for i := 0; i < *n; i++ {
		fmt.Scanln(&t[i].nama, &t[i].gold, &t[i].silver, &t[i].bronze)
		t[i].poin = poin(t[i].gold, t[i].silver, t[i].bronze)
	}
}

func poin(g, s, b int) int {
	return (g * 4) + (s * 3) + b
}

func sorting(t *olimpiade, n int) {
	var pass, i int
	var temp peserta
	pass = 1
	for pass <= n {
		i = pass
		temp = t[pass]
		for i > 0 && temp.poin > t[i-1].poin {
			t[i] = t[i-1]
			i--
		}
		for i > 0 && temp.poin == t[i-1].poin && temp.silver > t[i-1].silver {
			t[i] = t[i-1]
			i--
		}
		t[i] = temp
		pass++
	}
}

func tampilArray(t olimpiade, n int) {
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Println(t[i].nama, t[i].gold, t[i].silver, t[i].bronze)
	}
}

func main() {
	var tab olimpiade
	var n int
	isiArray(&tab, &n)
	sorting(&tab, n)
	tampilArray(tab, n)
}

// 5
// mpvh 8 4 8
// ptpc 2 7 10
// omen 8 9 5
// rmpw 2 8 7
// dnba 7 8 1
