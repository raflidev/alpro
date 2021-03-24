package main

import "fmt"

func main() {

	var n1, n2, key int

	fmt.Scanln(&key)
	for n1 != -1 && n2 != -1 {
		fmt.Scanln(&n1, &n2)

		if n1%key == 0 {
			fmt.Println(n2)
		} else if n2%key == 0 {
			fmt.Println(n1)
		}
	}
}
