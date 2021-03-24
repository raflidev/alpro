package main

import "fmt"

func main() {
	var nilai int
	fmt.Scan(&nilai)
	if nilai > 80 {
		if nilai == 100 {
			fmt.Println(100000)
		} else if nilai > 85 && nilai <= 99 {
			fmt.Println(50000)
		} else if nilai > 80 && nilai <= 85 {
			fmt.Println(25000)
		}
	} else {
		fmt.Println(0)
	}
}
