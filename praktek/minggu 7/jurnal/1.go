package main

import "fmt"

var biner string
var desimal int

func main() {
	fmt.Scanln(&desimal)
	biner = Des2Bin(desimal)
	fmt.Println(biner)
}

func Division(a, b int, result, remainder *int) {
	*result = a / b
	*remainder = a % b
}

func Num2Str(x int) string {
	if x == 0 {
		return "0"
	} else if x == 1 {
		return "1"
	} else {
		return ""
	}
}

func Des2Bin(desimal int) string {
	var hasil string
	var rDiv, rMod int
	hasil = ""
	for desimal > 0 {
		Division(desimal, 2, &rDiv, &rMod)
		desimal = rDiv
		hasil = Num2Str(rMod) + hasil
	}
	return hasil
}
