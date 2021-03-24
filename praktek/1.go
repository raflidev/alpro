package main

import "fmt"

func main() {
	var ma, mi, mo, n, jumlah, tahun int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Scanln(&mi, &mo, &ma)
		jumlah = jumlah + (mi - mo)
		if jumlah == ma {
			tahun = i
		}
	}
	if tahun != 0 {
		fmt.Println(tahun)
	} else {
		fmt.Println("data aman")
	}
}
