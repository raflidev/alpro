// package main

// import "fmt"

// const nMax = 37

// type tHimpunan struct {
// 	anggota [nMax]int
// 	panjang int
// }

// func bacaMasukan(h *tHimpunan) {
// 	var angka int
// 	fmt.Scan(&angka)
// 	i := 0
// 	for i < nMax && angka != -1 {
// 		if ada(*h, angka) == false {
// 			h.anggota[i] = angka
// 			h.panjang++
// 		}
// 		fmt.Scan(&angka)
// 		i++
// 	}
// }

// func ada(set tHimpunan, x int) bool {
// 	for i := 0; i < set.panjang; i++ {
// 		if set.anggota[i] == x {
// 			return true
// 		}
// 	}
// 	return false
// }

// func urut(set *tHimpunan) {
// 	var pass, i int
// 	var temp int
// 	pass = 1
// 	for pass <= set.panjang-1 {
// 		i = pass
// 		temp = set.anggota[pass]
// 		for i > 0 && temp < set.anggota[i-1] {
// 			set.anggota[i] = set.anggota[i-1]
// 			i--
// 		}
// 		set.anggota[i] = temp
// 		pass++
// 	}
// }

// func sama(h1, h2 tHimpunan) bool {
// 	return h1 == h2
// }

// func main() {
// 	var h1 tHimpunan
// 	var h2 tHimpunan
// 	fmt.Print("Anggota himpunan 1: ")
// 	bacaMasukan(&h1)
// 	fmt.Print("Anggota himpunan 2: ")
// 	bacaMasukan(&h2)
// 	urut(&h1)
// 	urut(&h2)
// 	fmt.Println("Himpunan 1 = Himpunan 2 ?", sama(h1, h2))
// }

// // 8 2 4 1 5 4 -1
// // 8 2 4 1 5 1 -1
// // 8 2 4 1 5 4 -1
// // 2 1 5 8 4 1 -1
// // 8 2 4 1 8 -1
// // 2 1 5 8 2 -1
