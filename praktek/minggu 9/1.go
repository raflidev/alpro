package main

import "fmt"

const nmax = 100

type tabGol = [nmax]int

func main() {
	var t tabGol
	var n int
	var t1, t2, t3 float64

	for i := 1; i <= 3; i++ {
		inputData(&t, &n)
		if i == 1 {
			t1 = rataan(&t, n)
		}
		if i == 2 {
			t2 = rataan(&t, n)
		}
		if i == 3 {
			t3 = rataan(&t, n)
		}
	}

	fmt.Println(t1, t2, t3)
}

func inputData(t *tabGol, n *int) {
	var data int
	data, *n = 0, 0
	fmt.Scan(&data)
	for data >= 0 && *n <= nmax {
		t[*n] = data
		fmt.Scan(&data)
		*n++
	}
}

func rataan(t *tabGol, n int) float64 {
	var sum int
	sum = 0
	for i := 0; i < n; i++ {
		sum += t[i]
	}

	return float64(sum / n)
}
