package main

import "fmt"

func main() {
	// kamus
	var tahun, mi, mo, ma, jumlah, salah int

	fmt.Scanln(&tahun)

	for i := 0; i < tahun; i++ {
		fmt.Scanln(&mi, &mo, &ma)
		jumlah = jumlah + (mi - mo)
		if jumlah == ma {
			salah = salah + 1
		}
	}
	if salah != tahun {
		fmt.Println(salah + 1)
	}
}
