package main

import "fmt"

func main() {
	var a, b, c, e, d, f rune

	fmt.Scanf("%c%c%c%c%c\n", &a, &b, &c, &d, &e)
	fmt.Scanln(&f)

	fmt.Printf("%c%c%c%c%c", a+f, b+f, c+f, d+f, e+f)
}
