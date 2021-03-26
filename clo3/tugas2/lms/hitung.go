package main

import "fmt"

var i int
var bilangan, rata, jumlah float64

func main() {
	fmt.Scan(&bilangan)
	i = 0
	rata = 0
	for bilangan > 0 {
		i = i + 1
		hitungRataRata(bilangan, i, &rata)
		fmt.Scan(&bilangan)
	}
	fmt.Println(rata)
}

func hitungRataRata(b float64, i int, r *float64) {
	jumlah = jumlah + b
	*r = jumlah / float64(i)
}
