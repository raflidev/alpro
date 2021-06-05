// package main

// import "fmt"

// const NMAX = 60

// type bunga struct {
// 	nama    string
// 	panjang int
// }

// type tabBunga = [NMAX]bunga

// func panjang(bunga string) int {
// 	var result int
// 	i := 0
// 	for string([]rune(bunga)[i]) != "." {
// 		bungatab := string([]rune(bunga)[i])
// 		if bungatab != "_" {
// 			result++
// 		}
// 		i++
// 	}
// 	return result

// }

// func isiArray(b *tabBunga, n *int) {
// 	var bunga string
// 	fmt.Scanln(*&n)
// 	for i := 0; i < *n; i++ {
// 		fmt.Scanln(&bunga)
// 		b[i].nama = bunga
// 		b[i].panjang = panjang(bunga)
// 	}
// }

// func mengurutkan(b *tabBunga, n int) {
// 	var pass, i int
// 	var temp bunga
// 	pass = 1
// 	for pass <= n-1 {
// 		i = pass
// 		temp = b[pass]
// 		for i > 0 && temp.panjang < b[i-1].panjang {
// 			b[i] = b[i-1]
// 			i--
// 		}
// 		b[i] = temp
// 		pass++
// 	}
// }

// func tampilArray(b tabBunga, n int) {
// 	for i := 0; i < n; i++ {
// 		fmt.Println(b[i].nama)
// 	}
// }

// func main() {
// 	var b tabBunga
// 	var n int
// 	isiArray(&b, &n)
// 	mengurutkan(&b, n)
// 	fmt.Println("")
// 	tampilArray(b, n)
// }

// // 7
// // Mawar.
// // Lili.
// // Kertas.
// // Forget_me_not.
// // Kamboja.
// // Anggrek.
// // Rafflesia.
