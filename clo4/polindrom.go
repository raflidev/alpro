package main

import "fmt"

func main() {
	var a = [5]int{10, 20, 5, 20, 10}
	fmt.Println(a)

	fmt.Println(a == reverse(a))
}

func reverse(x [5]int) [5]int {
	for i := 0; i < len(x)/2; i++ {
		j := len(x) - i - 1
		x[i], x[j] = x[j], x[i]
	}
	return x
}
