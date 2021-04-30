package main

import "fmt"

type buku struct {
	judul        string
	penulis      string
	tahun_terbit int
}

const nmax = 5

func main() {
	var tabBuku [nmax]buku
	var n int = 0
	tambahBuku(&tabBuku, &n)
	cariBuku(&tabBuku, &n, "C")
	tambahBuku(&tabBuku, &n)
	cariBuku(&tabBuku, &n, "KK")

}

func tambahBuku(kardus *[nmax]buku, atas *int) {
	var jumlah int
	fmt.Print("Mau input berapa buku?: ")
	fmt.Scanln(&jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Print("Judul: ")
		fmt.Scanln(&kardus[i].judul)
		fmt.Print("Penulis : ")
		fmt.Scanln(&kardus[i].penulis)
		fmt.Print("Tahun Terbit : ")
		fmt.Scanln(&kardus[i].tahun_terbit)
		*atas++
	}
}

func ambilBuku(kardus *[nmax]buku, atas *int, ambil *buku) {
	*ambil = kardus[*atas]
	*atas = *atas - 1
}

func cariBuku(kardus *[nmax]buku, atas *int, x string) {
	ketemu := -1
	n := 0
	var ambil buku
	for ketemu == -1 && n <= *atas {
		ambilBuku(kardus, atas, &ambil)
		if kardus[n].judul == x {
			ketemu = n
		} else {
			fmt.Print(kardus[n].judul, " ")
		}
		n++
	}

	if ketemu != -1 {
		fmt.Println("Ketemu")
	} else {
		fmt.Printf("\nBuku %s tidak ditemukan\n", x)
	}

}
