package main

import "fmt"

func main() {
	var x, counter int
	fmt.Scanln(&x)

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Print(i, " ")
			counter = counter + 1
		}
	}
	if counter == 2 {
		println("\nOlele")
	} else {
		println("\nBukan olele")
	}

}
