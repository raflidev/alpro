package main

import "fmt"

func main() {
	var i int
	var rata, bil int
	fmt.Scanln(&bil)
	i = 0
	rata = 0
	for bil > 0 {
		i = i + 1
		rata += bil
		fmt.Scanln(&bil)
	}
	fmt.Println(rata)
}
