package main

import "fmt"

func main() {
	var hour, minute, second, jumlah int

	fmt.Scanln(&hour, &minute, &second)

	jumlah = menghitungWaktu(hour, minute, second)
	fmt.Println(jumlah, "detik")
}

func menghitungWaktu(hour, minute, second int) int {
	return (hour * 3600) + (minute * 60) + (second * 1)
}
