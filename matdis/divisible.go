package main

import "fmt"

func main() {
	var count int
	for i := 1; i <= 100; i++ {
		if i%7 == 0 || i%8 == 0 {
			count += 1
		}
	}
	fmt.Println(count)
}
