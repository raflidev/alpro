package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&a, &b)
		if a%2 == 1 && b%2 == 0 {
			fmt.Println(b)
		} else if a%2 == 1 && b%2 == 1 {
			fmt.Println(a)
		} else if a%2 == 0 && b%2 == 0 {
			fmt.Println(a)
		} else if a%2 == 1 && b%2 == 1 {
			fmt.Println(a)
		}
	}
}
