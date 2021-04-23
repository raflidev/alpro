package main

import "fmt"

func main() {
	var a, b, hasil int
	fmt.Scanln(&a, &b)
	temp1 := a
	temp2 := b
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	hasil = (temp1 * temp2) / a
	fmt.Println(hasil)
}
