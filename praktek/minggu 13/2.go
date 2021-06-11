package main

import "fmt"

func main() {
	var t [100]int
	var n, u, d int
	fmt.Scanln(&n, &d, &u)
	isiArray(&t, &n)
	sorting(&t, u, d, n)
	tampilArray(t, n)

}

func isiArray(t *[100]int, n *int) {
	// input data berdasarkan n
	for i := 0; i < *n; i++ {
		fmt.Scan(&t[i])
	}
}

func tampilArray(t [100]int, n int) {
	// memberi jarak di hasil
	fmt.Println("")
	for i := 0; i < n; i++ {
		fmt.Print(t[i], " ")
	}
}

func sorting(t *[100]int, u, d, n int) {
	var pass, i int
	var temp int
	// memberikan jarak idx untuk sorting
	pass = d
	for pass <= u {
		i = pass
		temp = t[pass]
		for i > d && temp > t[i-1] {
			t[i] = t[i-1]
			i--
		}
		t[i] = temp
		pass++
	}
}

// 18 5 14
// 26 36 69 84 65 53 84 78 25 5 16 37 87 38 10 25 73 41
