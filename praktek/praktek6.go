package main

import "fmt"

func main() {
	var word string
	fmt.Scanln(&word)
	for word != "selesai" {
		fmt.Println(word)
		fmt.Scanln(&word)
	}
}
