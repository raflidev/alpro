package main

import "fmt"

func hitungVolume(r float64, t float64) float64 {
	var pi float64
	pi = 3.14
	return r * r * pi * t
}

func main() {
	var r, t float64
	fmt.Scanln(&r, &t)
	fmt.Println(hitungVolume(r, t))
}
