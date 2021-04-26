package main

import "fmt"

func main() {
	var tabInt [1080]int
	var jumlah, angka int
	i := 0
	for true {
		fmt.Scan(&tabInt[i])
		if tabInt[i] == -1 {
			break
		}
		jumlah = i
		i = i + 1
	}
	fmt.Println(tabInt)
	fmt.Scanln(&angka)
	fmt.Println(frekuensi(tabInt, jumlah, angka))

}

func frekuensi(t [1080]int, n, x int) int {
	var count int
	i := 0
	for i < n {
		if x == t[i] {
			count = count + 1
		}
		i = i + 1
	}
	return count
}
