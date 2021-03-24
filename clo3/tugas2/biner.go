package main

import (
	"fmt"
	"math"
)

func main() {
	var bil int
	fmt.Scanln(&bil)
	biner(float64(bil))
}

func biner(bil float64) {
	if bil >= 1 {

		biner(math.Floor(float64(bil) / float64(2)))
	}
	fmt.Print(int(bil)%2, " ")
}
