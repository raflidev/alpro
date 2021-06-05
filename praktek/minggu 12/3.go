package main

import "fmt"

const Nmax int = 1000000

type tabDate [Nmax]string

func addData(list *tabDate, n *int) {
	var tanggal string
	fmt.Scanln(&tanggal)
	for tanggal != "####.##.##" {
		list[*n] = tanggal
		*n++
		fmt.Scanln(&tanggal)
	}
}

func urut(list *tabDate, n int) {
	var (
		i, idx int
	)
	for i < Nmax {
		idx = maxPos(*list, i, n)
		swap(&list[i], &list[idx])
	}
}

func maxPos(list tabDate, start, n int) int {
	var (
		max    string
		result int
	)
	max = list[start]
	result = start
	start++
	for start < n {
		if list[start] > max {
			max = list[start]
			result = start
		}
		start++
	}
	return result
}

func swap(x, y *string) {
	var tmp string
	tmp = *x
	*x = *y
	*y = tmp
}

func printData(list tabDate, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(list[i])
	}
}

func main() {
	var t tabDate
	var n int
	addData(&t, &n)
	urut(&t, n)
	printData(t, n)
}
