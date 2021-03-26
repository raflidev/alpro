package main

import "fmt"

func main() {
	var op, op1, op2 string
	fmt.Scanln(&op, &op1, &op2)
	tampil(op, op1, op2)
}

func nilaiTeks(n string) int {
	var x int
	if n == "NOL" {
		x = 0
	} else if n == "SATU" {
		x = 1
	} else if n == "DUA" {
		x = 2
	} else if n == "TIGA" {
		x = 3
	} else if n == "EMPAT" {
		x = 4
	} else if n == "LIMA" {
		x = 5
	} else if n == "ENAM" {
		x = 6
	} else if n == "TUJUH" {
		x = 7
	} else if n == "DELAPAN" {
		x = 8
	} else if n == "SEMBILAN" {
		x = 9
	}
	return x
}
func teksNilai(n int) string {
	var x string
	if n == 0 {
		x = "NOL"
	} else if n == 1 {
		x = "SATU"
	} else if n == 2 {
		x = "DUA"
	} else if n == 3 {
		x = "TIGA"
	} else if n == 4 {
		x = "EMPAT"
	} else if n == 5 {
		x = "LIMA"
	} else if n == 6 {
		x = "ENAM"
	} else if n == 7 {
		x = "TUJUH"
	} else if n == 8 {
		x = "DELAPAN"
	} else if n == 9 {
		x = "SEMBILAN"
	}
	return x
}

func hasilKomputasi(op, op1, op2 string) int {
	var jumlah int
	if op == "TAMBAH" {
		jumlah = nilaiTeks(op1) + nilaiTeks(op2)
	} else if op == "KURANG" {
		jumlah = nilaiTeks(op1) - nilaiTeks(op2)
	}
	return jumlah
}

func tampil(op, op1, op2 string) {
	var result int
	result = hasilKomputasi(op, op1, op2)
	if result < 0 {
		result = result * -1
		fmt.Println("MINUS NOL", teksNilai(result))
	} else if result > 0 && result < 9 {
		fmt.Println("NOL", teksNilai(result))
	} else if result >= 10 {
		result = result - 10
		fmt.Println("SATU", teksNilai(result))
	}
}
