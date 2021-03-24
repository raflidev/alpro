package main

import "fmt"

func main() {
	var i, count int
	fmt.Scanln(&i)
	for x := 0; x < i; x++ {
		count = count + 2
		fmt.Print(count, " ")
	}
}
