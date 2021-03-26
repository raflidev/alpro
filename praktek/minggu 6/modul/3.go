package main

import (
	"fmt"
	"math"
)

func main() {
	var cx, cy, r, x, y float64
	fmt.Scanln(&cx, &cy, &r)
	fmt.Scanln(&x, &y)
	if posisi(cx, cy, r, x, y) {
		fmt.Println("anda berada didalam taman")
	} else {
		fmt.Println("anda berada diluar taman")
	}
}

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func posisi(cx, cy, r, x, y float64) bool {
	var jarak1 float64
	jarak1 = jarak(cx, cy, x, y)

	return jarak1 < r
}
