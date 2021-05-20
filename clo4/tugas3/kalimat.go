package main

import "fmt"

type tabStr = [100]string

func panjang(s string) int {
	// pseudocode
	// var x int
	// x = 0
	// while s[x] !=  null{
	// 	x++
	// }
	// return x
	return len(s)
}

func isiArray(s string, daftar *tabStr, k *int) {
	for i := 0; i < panjang(s); i++ {
		if string([]rune(s)[i]) != "_" {
			daftar[i] = string([]rune(s)[i])
			*k++
		}
	}

	// pseudocode
	// var i int
	// i = 0
	// for i to panjang(s) do
	// 	if s[i] != "_"
	// 		daftar[i] = s[i]
	// 		k = k+1
	// 	endif
	// endfor
}

func cariIndex(daftar tabStr, k int, w string, idx *int) {
	found := -1
	i := 0
	for found == -1 && i <= k {
		if daftar[i] == string([]rune(w)[0]) {
			found = i
		}
		i++
	}

	*idx = found
}

// abcd_efgh efgh
// qwerty_asdfg rty
// purwadjlcn oke

func cariPosisi(daftar *tabStr, n int, w string, posisi, idx *int) {
	spasi := 0
	for i := 0; i <= n; i++ {
		if daftar[i] == "" {
			spasi = i
		}
	}

	if spasi == *idx || *idx > spasi {
		*posisi = 2
	} else {
		*posisi = 1
	}
	if spasi == n {
		*idx = -1
	} else if *idx > spasi {
		if (*idx - (spasi + 1)) == 0 {
			*idx = *idx - (spasi + 1)
		} else {
			*idx = (*idx - 1) - spasi
		}
	} else {
		if ((spasi + 1) - *idx) == 0 {
			*idx = (spasi + 1) - *idx
		} else {
			*idx = spasi - *idx - 1
		}

	}
}

func main() {
	var t tabStr
	var teks, w string
	var k, idx, posisi int

	fmt.Scanln(&teks, &w)

	k = 0
	idx = 0
	posisi = 0
	isiArray(teks, &t, &k)
	cariIndex(t, k, w, &idx)
	cariPosisi(&t, k, w, &posisi, &idx)

	if idx == -1 {
		fmt.Println(idx)
	} else {
		fmt.Println(posisi, idx)
	}
}
