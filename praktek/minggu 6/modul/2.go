package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scanln(&x)
	fmt.Println("f(", x, ") = ", f(x))
	fmt.Println("g(", x, ") = ", g(x))
	fmt.Println("h(", x, ") = ", h(x))
	komposisi(x, &y)
	fmt.Println("(fogoh)(", x, ") =", y)
}

func komposisi(x float64, y *float64) {
	*y = f(g(h(x)))
}

func f(x float64) float64 {
	return x * x
}
func g(x float64) float64 {
	return x - 2
}
func h(x float64) float64 {
	return x + 1
}
