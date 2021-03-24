package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5 int
	var s1, s2, s3, s4, s5 int

	var quads bool
	fmt.Scanf("%d%c %d%c %d%c %d%c %d%c", &a1, &s1, &a2, &s2, &a3, &s3, &a4, &s4, &a5, &s5)
	quads = a2 == a3 && a3 == a4 && a4 == a5 || a1 == a3 && a3 == a4 && a4 == a5 || a1 == a2 && a2 == a4 && a4 == a5 || a1 == a2 && a2 == a3 && a3 == a5 || a1 == a2 && a2 == a3 && a3 == a4
	println(quads)
}
