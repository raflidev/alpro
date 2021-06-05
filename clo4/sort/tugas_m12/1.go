// package main

// import "fmt"

// const nmax = 60

// type tabAngka = [nmax]float64

// func ganjilAtauGenap(baris int) string {
// 	if baris%2 == 0 {
// 		return "genap"
// 	}
// 	return "ganjil"
// }

// func sorting(A *tabAngka, n int) {
// 	var pass, i int
// 	var temp float64
// 	pass = 1
// 	for pass <= n-1 {
// 		i = pass
// 		temp = A[pass]
// 		for i > 0 && temp < A[i-1] {
// 			A[i] = A[i-1]
// 			i--
// 		}
// 		A[i] = temp
// 		pass++
// 	}
// }

// func cariMedian(angka tabAngka, baris int) {
// 	if ganjilAtauGenap(baris) == "ganjil" {
// 		fmt.Println(angka[(baris / 2)])
// 	} else {
// 		fmt.Println((angka[(baris/2)] + angka[(baris/2)-1]) / 2)
// 	}
// }

// func main() {
// 	var baris int
// 	var angka tabAngka
// 	fmt.Scanln(&baris)

// 	for i := 0; i < baris; i++ {
// 		fmt.Scan(&angka[i])
// 	}

// 	sorting(&angka, baris)
// 	fmt.Println(angka)
// 	cariMedian(angka, baris)
// }
