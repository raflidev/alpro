package main

import "fmt"

const nmax int = 2022

type student struct {
	name, sid string
	gpa       float64
}
type tabMhs = [nmax - 1]student

func main() {
	var tabmah tabMhs
	var n int
	inputArray(&tabmah, &n)
	sortArray(&tabmah, n)
	printArray(tabmah, n)

}

func inputArray(T *tabMhs, n *int) {
	for i := 0; i < nmax-1; i++ {
		fmt.Scanln(&T[i].sid, &T[i].name, &T[i].gpa)
		if T[i].sid == "#" {
			break
		}
	}
}

func sortArray(T *tabMhs, n int) {
	var pass, idx, i int
	var temp student
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if T[idx].gpa < T[i].gpa {
				idx = i
			}
			i++
		}

		// T[idx], T[pass-1] = T[pass-1], T[idx]
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
		pass = pass + 1
	}
}

func printArray(T tabMhs, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i].sid, T[i].name, T[i].gpa)
	}
}

// A 15 70.4
// B 11 75.1
// C 12 30.9
// D 13 98.1
// E 16 91.1
// F 17 91.3
// G 18 91.2
// H 19 91.6
// I 20 91.8
// J 21 91.7
// #
