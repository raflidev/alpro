package main

import "fmt"

const nmax = 100

type tabAngka = [nmax]float64

func main() {
	var t tabAngka
	var n float64

	i := 0
	fmt.Scan(&t[i])
	for i < nmax && t[i] != -1 {
		i++
		n++
		fmt.Scan(&t[i])
	}

	fmt.Println(rata2(t, 0, n))
}

func rata2(t tabAngka, i int, n float64) float64 {
	if float64(i) == n-1 {
		return t[i]
	}

	if i == 0 {
		return (t[i] + rata2(t, i+1, n)) / n
	}

	return t[i] + rata2(t, i+1, n)
}
