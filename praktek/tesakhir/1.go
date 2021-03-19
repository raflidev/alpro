package main

import (
	"fmt"
	"math"
)

func main() {
	var bil1, bil2 int
	var rusak int
	var selisih float64
	var jumlah int
	for true {
		fmt.Scanln(&bil1, &bil2)
		jumlah = bil1 + bil2
		if bil1 > 0 && bil2 > 0 || bil1 < 0 && bil2 < 0 {
			rusak = rusak + 1
		} else if jumlah == 0 {
			break
		} else {
			selisih = selisih + math.Abs(float64(bil1)+float64(bil2))
		}

	}
	fmt.Println("Terminal rusak:", rusak)
	fmt.Println("Total selisih arus:", selisih)

}
