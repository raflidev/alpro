package main

import "fmt"

func main() {
	var angka int
	fmt.Scanln(&angka)
	fmt.Println(factorial(angka))
}

func factorial(n int) int {
	if n > 1 {
		fmt.Println(n - 1)
		return n * factorial(n-1)
	} else {
		return 1
	}
}
