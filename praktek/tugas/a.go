package main

import "fmt"

func main() {
	var i int
	i = 10000
	for x := 1; x < i; x++ {
		if x%2 == 0 && x%5 == 0 {
			fmt.Print(x, " ")
		}
	}
}
