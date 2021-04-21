package main

import "fmt"

type tim struct {
	tim1, tim2, tim3 int
}

func main() {
	var tabGol []tim
	var n int
	var j1, j2, j3 float64

	fmt.Print("Berapa kemenangan? : ")
	fmt.Scanln(&n)
	inputData(&tabGol, n)

	j1, j2, j3 = rataan(tabGol, n)
	fmt.Printf("Rata-Rata Kemenangan Tim 1: %g \n", j1)
	fmt.Printf("Rata-Rata Kemenangan Tim 2: %g \n", j2)
	fmt.Printf("Rata-Rata Kemenangan Tim 3: %g \n", j3)

}

func inputData(tabGol *[]tim, n int) {
	a := make([]tim, n)
	i := 0
	for i < n {
		fmt.Scanln(&a[i].tim1, &a[i].tim2, &a[i].tim3)
		i++
	}
	*tabGol = a
}

func rataan(tabGol []tim, n int) (float64, float64, float64) {
	var jumlah_t1, jumlah_t2, jumlah_t3 int
	i := 0
	for i < n {
		jumlah_t1 = jumlah_t1 + tabGol[i].tim1
		jumlah_t2 = jumlah_t2 + tabGol[i].tim2
		jumlah_t3 = jumlah_t3 + tabGol[i].tim3
		i++
	}
	return float64(jumlah_t1 / n), float64(jumlah_t2 / n), float64(jumlah_t3 / n)
}
