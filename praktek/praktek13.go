package main

import "fmt"

func main() {
	var i string
	fmt.Scanln(&i)

	for k := 0; k < len(i); k++ {
		fmt.Println(string(i[2 : k+2]))
	}
}
