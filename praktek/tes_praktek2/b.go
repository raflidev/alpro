package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&a, &b)
		if a%2 == 0 && b%2 == 1 {
			println(a)
		} else if a%2 == 0 && b%2 == 0 {
			println(a)
		} else if a%2 == 1 && b%2 == 0 {
			println(b)
		} else if a%2 == 1 && b%2 == 1 {
			println(a)
		}
	}

}
