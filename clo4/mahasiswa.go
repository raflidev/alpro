package main

import "fmt"

type Mhs struct {
	nama                     string
	nim, eprt, ipk, semester int
}

func main() {
	var Mahasiswa [1000]Mhs
	var n, jumlah int
	n = 0
	for true {
		fmt.Scan(&Mahasiswa[n].nama, &Mahasiswa[n].nim, &Mahasiswa[n].eprt, &Mahasiswa[n].semester, &Mahasiswa[n].ipk)
		if Mahasiswa[n].nama == "none" {
			break
		}
		jumlah = n
		n = n + 1
	}

	fmt.Printf("EPRT tertinggi : %d \n", cariEprt(Mahasiswa, jumlah))
	fmt.Printf("IPK terrendah : %d \n", cariIpk(Mahasiswa, jumlah))
	fmt.Printf("Rata-rata semester lulusan : %d \n", cariSemester(Mahasiswa, jumlah))
}

func cariEprt(Mahasiswa [1000]Mhs, jumlah int) int {
	max := 0
	i := 0
	for i < jumlah {
		if Mahasiswa[max].eprt < Mahasiswa[i].eprt {
			max = Mahasiswa[i].eprt
		}
		i = i + 1
	}
	return max
}

func cariIpk(Mahasiswa [1000]Mhs, jumlah int) int {
	min := Mahasiswa[0].eprt
	i := 0
	for i < jumlah {
		if min > Mahasiswa[i].eprt {
			min = Mahasiswa[i].eprt
		}
		i = i + 1
	}
	return min
}

func cariSemester(Mahasiswa [1000]Mhs, jumlah int) int {
	tambah := 0
	i := 0
	for i <= jumlah {
		tambah = tambah + Mahasiswa[i].semester
		i = i + 1
	}
	return tambah / (jumlah + 1)
}
