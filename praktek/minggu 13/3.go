package main

import "fmt"

const nmax = 100

type belanja = [nmax]int

func main() {
	var t belanja
	var n, ha int
	var hp float64
	// masukan dan menghitung perbelanjaan
	entri(&t, &n)

	ha = total(t, n)
	if ha > 100000 {
		// sorting
		urut(&t, n)
		// promo
		hp = 0
		i := 0
		for i < n {
			// menghitung promo dan meng-sum nya
			hp += float64(t[i+1] - (t[i+1] * (i + 1) / 100))
			i++
		}
		// cetak keluaran
		fmt.Println(ha, hp)
	} else {
		// cetak keluaran
		fmt.Println(ha, ha)
	}
}

func entri(t *belanja, n *int) {
	var harga, satuan int
	i := 0
	fmt.Scanln(&harga, &satuan)
	// berhenti jika harga dan satuan 0
	for i < nmax && harga != 0 && satuan != 0 {
		// mengkalkulasi harga belanja
		t[i] = harga * satuan
		i++
		// menambah n arrau
		*n++
		fmt.Scanln(&harga, &satuan)
	}
}

func urut(t *belanja, n int) {
	// sorting
	var pass, i int
	var temp int
	pass = 1
	for pass <= n {
		i = pass
		temp = t[pass]
		// sorting asc
		for i > 0 && temp < t[i-1] {
			t[i] = t[i-1]
			i--
		}
		t[i] = temp
		pass++
	}
}

func total(t belanja, n int) int {
	var sum int
	i := 0
	for i < n {
		// sum total
		sum = sum + t[i]
		i++
	}
	return sum
}

// 16800 2
// 12100 3
// 9100 1
// 18300 2
// 13300 1
// 11300 3
// 0 0
