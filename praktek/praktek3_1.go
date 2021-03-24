package main

import "fmt"

func main() {
	var a, b, c, d, e, x, y, z byte
	fmt.Scanln(&a, &b, &c, &d, &e)
	fmt.Scanf("%c%c%c", &x, &y, &z)
	fmt.Printf("%c%c%c%c%c\n%c%c%c", a, b, c, d, e, x+1, y+1, z+1)
}
