package main

import "fmt"

var uangAwal float64
var tUang, tPengeluaran, tPemasukan float64

func main() {

	fmt.Scan(&uangAwal)
	tUang = uangAwal

	membeliBibit(uangAwal, &tUang, &tPengeluaran)
	membeliPupuk(uangAwal, &tUang, &tPengeluaran)
	menanamBibit(uangAwal, &tUang, &tPengeluaran)
	menanamBibit(uangAwal, &tUang, &tPengeluaran)
	memberiPupuk(uangAwal, &tUang, &tPengeluaran)
	memberiPupuk(uangAwal, &tUang, &tPengeluaran)
	memberiPupuk(uangAwal, &tUang, &tPengeluaran)
	memberiPupuk(uangAwal, &tUang, &tPengeluaran)
	memanen(uangAwal, &tUang, &tPengeluaran)
	memanen(uangAwal, &tUang, &tPengeluaran)
	menjualHasilPanen(uangAwal, &tUang, &tPengeluaran, &tPemasukan)

	fmt.Printf("%.2f %.2f %.2f\n", tPengeluaran, tPemasukan, tUang)
}

func membeliBibit(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = *tUang - uangAwal/4
	*tPengeluaran = uangAwal / 4
}

func membeliPupuk(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = *tUang - uangAwal/20
	*tPengeluaran = *tPengeluaran + uangAwal/20
}

func menanamBibit(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = *tUang - uangAwal/200
	*tPengeluaran = *tPengeluaran + uangAwal/200
}

func memberiPupuk(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = *tUang - uangAwal/200
	*tPengeluaran = *tPengeluaran + uangAwal/200
}

func memanen(uangAwal float64, tUang, tPengeluaran *float64) {
	*tUang = *tUang - 3*uangAwal/200
	*tPengeluaran = *tPengeluaran + 3*uangAwal/200
}

func menjualHasilPanen(uangAwal float64, tUang, tPengeluaran, tPemasukan *float64) {
	*tUang = *tUang - 3*uangAwal/200 + uangAwal/2
	*tPengeluaran = *tPengeluaran + 3*uangAwal/200
	*tPemasukan = uangAwal / 2
}
