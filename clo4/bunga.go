package main

import "fmt"

func main() {
	var bunga [100]string
	var n, jumlah int
	var hapusNama string
	n = 0
	fmt.Println("Isi nama bunga, kosongkan jika sudah selesai")
	for true {
		fmt.Scanln(&bunga[n])
		if bunga[n] == "" {
			break
		}
		jumlah = n
		n = n + 1
	}
	fmt.Println("Nama yang ingin dihapus")
	fmt.Scanln(&hapusNama)

	if hapusNamaBunga(bunga, hapusNama, jumlah) == -1 {
		fmt.Println("bunga tidak ditemukan")
	} else {
		fmt.Println("Sebelum dihapus")
		fmt.Println(bunga)
		bunga[hapusNamaBunga(bunga, hapusNama, jumlah)] = ""
		fmt.Println("Sesudah dihapus")
		fmt.Println(bunga)
	}

}

func hapusNamaBunga(bunga [100]string, hapusNama string, jumlah int) int {
	ketemu := -1
	k := 0
	for ketemu == -1 && k <= jumlah {
		if bunga[k] == hapusNama {
			ketemu = k
		}
		k = k + 1
	}
	return ketemu
}
