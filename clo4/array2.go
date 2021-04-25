package main

import "fmt"

func main() {
	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(reverse1(array))
}

func reverse1(x [5]int) [5]int {
	for i := 0; i < len(x)/2; i++ {
		j := len(x) - i - 1
		x[i], x[j] = x[j], x[i]
	}
	return x
}
