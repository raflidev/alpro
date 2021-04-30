package main

import (
	"fmt"
)

const nmax = 127

type table = [nmax]string

func main() {
	var t table
	var n int
	n = 0
	isiArray(&t, &n)
	balikanArray(&t, n)
	cetakArray(t, n)
	fmt.Println(polindrom(t, n))

}

func isiArray(t *table, n *int) {
	var data string
	fmt.Scan(&data)
	for data != "." && *n <= nmax {
		t[*n] = data
		fmt.Scan(&data)
		*n++
	}
}

func cetakArray(t table, n int) {
	for i := 0; i <= n; i++ {
		fmt.Print(t[i], " ")
	}
}

func balikanArray(t *table, n int) {
	for i := 0; i <= (n-1)/2; i++ {
		temp := t[i]
		t[i] = t[(n-1)-i]
		t[(n-1)-i] = temp
	}
}

func polindrom(t table, n int) bool {
	t2 := t
	balikanArray(&t2, n)
	return t == t2
}
