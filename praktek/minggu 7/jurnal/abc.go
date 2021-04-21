package main

import "fmt"

var a, b int

func main() {
	fmt.Scanln(&a, &b)
	for a <= b {
		abc(&a, b)
		b = b - a
	}
	fmt.Println(a, b)
}
func abc(x *int, y int) {
	fmt.Println(*x, y)
	*x = *x + 1
	y = y - 1
}
