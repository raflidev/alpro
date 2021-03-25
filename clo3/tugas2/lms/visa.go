package main

import "fmt"

var tanggal, tahun int
var bulan, hari string
var tgl2, bln2, thn2 int

func main() {
	fmt.Scanln(&tanggal, &bulan, &tahun, &hari)
	Pengembalian(tanggal, Bulan2Angka(bulan), tahun, hari, &tgl2, &bln2, &thn2)
	fmt.Println(tgl2, Angka2Bulan(bln2), thn2)
}

func Kabisat(tahun int) bool {
	if tahun%4 == 0 {
		if tahun%100 == 0 {
			if tahun%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func Bulan2Angka(bulan string) int {
	if bulan == "januari" {
		return 1
	} else if bulan == "februari" {
		return 2
	} else if bulan == "maret" {
		return 3
	} else if bulan == "april" {
		return 4
	} else if bulan == "mei" {
		return 5
	} else if bulan == "juni" {
		return 6
	} else if bulan == "juli" {
		return 7
	} else if bulan == "agustus" {
		return 8
	} else if bulan == "september" {
		return 9
	} else if bulan == "oktober" {
		return 10
	} else if bulan == "november" {
		return 11
	} else if bulan == "desember" {
		return 12
	} else {
		return 0
	}
}
func Angka2Bulan(angka int) string {
	if angka == 1 {
		return "januari"
	} else if angka == 2 {
		return "februari"
	} else if angka == 3 {
		return "maret"
	} else if angka == 4 {
		return "april"
	} else if angka == 5 {
		return "mei"
	} else if angka == 6 {
		return "juni"
	} else if angka == 7 {
		return "juli"
	} else if angka == 8 {
		return "agustus"
	} else if angka == 9 {
		return "september"
	} else if angka == 10 {
		return "oktober"
	} else if angka == 11 {
		return "november"
	} else if angka == 12 {
		return "desember"
	} else {
		return ""
	}
}

func JumlahHari(bln, thn int) int {
	var hari int

	if Kabisat(thn) && bln == 2 {
		hari = 29
	} else if bln == 2 {
		hari = 28
	} else if (bln%2 == 0 && bln < 8) || (bln%2 == 1 && bln > 8) {
		hari = 30
	} else {
		hari = 31
	}

	return hari
}

func Pengembalian(tgl1, bln1, thn1 int, hari string, tgl2, bln2, thn2 *int) {
	var jHari int
	jHari = JumlahHari(bln1, thn1)
	if hari == "kamis" || hari == "jumat" {
		tgl1 += 4
	} else {
		tgl1 += 2
	}
	if tgl1 > jHari {
		tgl1 = (tgl1) - jHari
		if bln1 == 12 {
			bln1 = 1
			thn1 = thn1 + 1
		} else {
			bln1 = bln1 + 1
		}
	}

	*tgl2 = tgl1
	*bln2 = bln1
	*thn2 = thn1
}
