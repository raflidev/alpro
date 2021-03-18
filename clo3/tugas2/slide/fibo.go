package main

import "fmt"

func main() {
	var n, t1, t2, nth, jumlah int
	t1, t2 = 0, 1
	fmt.Scanln(&n)

	jumlah = fibo(jumlah, t1, t2, nth, n)

	fmt.Println("Jumlah fibo:", jumlah)
}

func fibo(jumlah, t1, t2, nth, n int) int {
	for i := 0; i < n; i++ {
		jumlah = jumlah + t1
		nth = t1 + t2
		t1 = t2
		t2 = nth
	}
	return jumlah
}
