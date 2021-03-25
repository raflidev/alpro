package main

import "fmt"

var member string
var a, b, c, d, e int
var cashback, diskon float64

func main() {
	masukan()
	promoCashbackDiskon()
	fmt.Println(cashback, diskon)
}

func masukan() {
	fmt.Scanln(&member, &a, &b, &c, &d, &e)
}

func cekGanjil() bool {
	if a%2 == 1 && b%2 == 1 && c%2 == 1 && d%2 == 1 && e%2 == 1 {
		return true
	} else {
		return false
	}
}

func cekGenap() bool {
	if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		return true
	} else {
		return false
	}
}

func promoCashbackDiskon() {
	if cekGanjil() {
		diskon = 1.7 * float64(c+d+e)
	} else {
		cashback = 3.1 * float64(a+b+c)
	}
	if !cekGanjil() && !cekGenap() {
		cashback = (3.1 * float64(a+b+c))
		diskon = (1.7 * float64(c+d+e))
	}
	membership()

	if cashback > 35 {
		cashback = 35
	}
	if diskon > 50 {
		diskon = 50
	}
}

func membership() {
	var memberdiskon, membercashback float64
	if member == "yes" {
		memberdiskon = diskon * 15 / 100
		membercashback = cashback * 15 / 100
		diskon = diskon + memberdiskon
		cashback = cashback + membercashback
	}
}
