package main

import "fmt"

func main() {
	var i int
	fmt.Scanln(&i)

	fmt.Println(i / 1000)
	fmt.Println(i / 100 % 10)
	fmt.Println(i / 10 % 10)
	fmt.Println(i % 10)

	fmt.Println(i / 100)
	fmt.Println(i / 10 % 100)
	fmt.Println(i % 100)

	fmt.Println(i / 10)
	fmt.Println(i % 1000)
}
