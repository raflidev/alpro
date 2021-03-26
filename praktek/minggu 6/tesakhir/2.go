package main

import "fmt"

func main() {
	var n int
	var p, q, r string
	fmt.Scanln(&n, &p, &q, &r)
	// percabangan fungsi
	if p == "f" && q == "g" && r == "h" {
		fmt.Println(h(g(f(n))))
	} else if p == "g" && q == "h" && r == "f" {
		fmt.Println(f(h(g(n))))
	} else if p == "h" && q == "f" && r == "g" {
		fmt.Println(g(f(h(n))))
	} else if p == "h" && q == "g" && r == "f" {
		fmt.Println(f(g(h(n))))
	} else if p == "g" && q == "f" && r == "h" {
		fmt.Println(h(f(g(n))))
	} else if p == "f" && q == "h" && r == "g" {
		fmt.Println(g(h(f(n))))
	} else if p == "f" && q == "f" && r == "f" {
		fmt.Println(f(f(f(n))))
	}
}

// fungsi f g dan h
func f(x int) int {
	return (2 * x) + 5
}
func g(x int) int {
	return (x * x) + (2 * x)
}
func h(x int) int {
	return x - 3
}
