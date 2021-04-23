package main

import "fmt"

func main() {
	var a, b, c int
	var no string
	no = ""
	for no != "." {
		fmt.Scan(&no)
		hitungNopol(no, &a, &b, &c)
	}
	fmt.Println(a, b, c)
}
func hitungNopol(noPol string, nBandung, nJakarta, nLainnya *int) {
	if noPol == "D" || noPol == "Z" {
		*nBandung += 1
	} else if noPol == "B" || noPol == "F" {
		*nJakarta += 1
	} else if noPol != "." {
		*nLainnya += 1
	}
}
