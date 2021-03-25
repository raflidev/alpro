package main

import (
	"fmt"
)

var a, b int
var remainder, result int
var bucket string

func main() {
	fmt.Scanln(&a)
	b = 2
	aa := Des2Bin(a)
	fmt.Println(aa)
}
func Division(a, b int, result *int, remainder int) {
	for a != 0 {
		if a%2 == 0 {
			remainder = 0
		} else {
			remainder = 1
		}
		bucket = Num2Str(remainder) + bucket
		a /= b
	}
}
func Num2Str(x int) string {
	if x == 0 {
		return "0"
	} else {
		return "1"
	}
}

func Des2Bin(desimal int) string {
	Division(a, b, &result, remainder)
	return bucket
}
