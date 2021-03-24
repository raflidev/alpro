package main

import "fmt"

func main() {
	var x, total int
	fmt.Scanln(&x)
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Print(i, " ")
			total = total + 1
		}
	}
	if total == 2 {
		print("\nOleole")
	} else {
		print("\nBukan oleole")
	}
}
