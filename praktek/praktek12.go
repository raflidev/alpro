package main

import "fmt"

func main() {
	var i string
	var x int
	fmt.Scanln(&i, &x)

	if x >= -27 && x <= 27 {
		for k := 0; k < len(i); k++ {
			hasil := int(i[k]) + x
			fmt.Printf("%c", hasil)
		}
	}

}
