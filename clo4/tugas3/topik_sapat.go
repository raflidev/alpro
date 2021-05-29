package main

import "fmt"

type tabStr = [100]string

func main() {
	var t tabStr
	var teks, w string
	var k, idx, posisi int
	fmt.Scanln(&teks, &w)
	isiArray(teks, &t, &k)
	cariIndex(&t, &k, w, &idx)
	if idx == -1 {
		fmt.Println(-1)
	} else {
		cariPosisi(t, k, w, &posisi, &idx)
		fmt.Println(posisi, idx)
	}
}

func panjang(s string) int {
	return len(s)
}

func isiArray(s string, daftar *tabStr, k *int) {
	for *k < panjang(s) {
		daftar[*k] = string(s[*k])
		*k++
	}
}

func cariIndex(daftar *tabStr, k *int, w string, idx *int) {
	var i int
	var ketemu bool
	for i < *k && !ketemu {
		if daftar[i] == string(w[0]) {
			ketemu = true
			*idx = i
		}
		i++
	}
	if !ketemu {
		*idx = -1
	}
}

func cariPosisi(daftar tabStr, n int, w string, posisi *int, idx *int) {
	var i int
	*posisi = 1
	for i < n && daftar[i] != daftar[*idx] {
		if daftar[i] == "_" {
			*posisi += 1
			*idx -= i + 1
		}
		i++
	}
}
