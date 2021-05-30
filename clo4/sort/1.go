package main

import "fmt"

const nmax = 20

type array = [nmax]int

func mengurutkan(A *array, n int) {
	var pass, idx, i, temp int
	pass = 1
	for pass < n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx] < A[i] {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func main() {
	var angka array
	var n, data int
	i := 0
	fmt.Scan(data)
	for i < nmax && data != -1 {
		angka[i] = data
		fmt.Scan(&data)
		n = n + 1
		i = i + 1
	}

	fmt.Println(angka)
	mengurutkan(&angka, n)
	fmt.Println(angka)
}

// 12 3 45 67 8 90 679 23 42 56 7 89 30 126 14 5 2 90 -1
