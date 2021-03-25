package main

import "fmt"

func main() {
	var a, b, c float64
	var op string

	fmt.Scanln(&a, &b, &c, &op)

	if op == "T" {
		fmt.Println(a * b * c)
	} else if op == "O" {
		fmt.Println(3.14 * a * a * c)
	} else if op == "P" {
		fmt.Println(0.5 * a * b * c)
	}
}
