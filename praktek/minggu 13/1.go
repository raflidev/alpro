package main

import (
	"fmt"
)

const nmax = 100

type mahasiswa struct {
	nama   string
	tinggi int
}

type dataMhs = [nmax]mahasiswa

func main() {
	var m dataMhs
	var n int
	bacaData(&m, &n)
	urutData(&m, n)
	cetakData(m, n)
}

func bacaData(t *dataMhs, n *int) {
	// input data
	fmt.Scanln(&*n)
	for i := 0; i < *n; i++ {
		fmt.Scanln(&t[i].nama, &t[i].tinggi)
	}
}

func cetakData(t dataMhs, n int) {
	// cetak jarak/row kosong
	fmt.Println("")

	// cetak data
	for i := 0; i < n; i++ {
		fmt.Println(t[i].nama, t[i].tinggi)
	}
}

func urutData(t *dataMhs, n int) {
	var pass, idx, i int
	var temp mahasiswa
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			// asc
			if t[idx].tinggi > t[i].tinggi {
				idx = i
			}
			i++
		}
		// tuker tempat index
		temp = t[pass-1]
		t[pass-1] = t[idx]
		t[idx] = temp
		pass++
	}
}

// 18
// BN 170
// SA 161
// OP 163
// UZ 160
// BQ 170
// KC 162
// XL 156
// NE 174
// LM 168
// JV 163
// OF 176
// JV 166
// EO 179
// HO 174
// MB 173
// AH 159
// GJ 172
// LY 162
