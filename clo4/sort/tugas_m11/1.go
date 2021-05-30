package main

import "fmt"

const nmax = 2022

type student struct {
	name, sid string
	gpa       float64
}

type tabMhs = [nmax - 1]student

func input(T *tabMhs, n *int) {
	var i int
	var name, sid string
	var gpa float64
	fmt.Scanln(&name, &sid, &gpa)
	for i < nmax && name != "#" {
		T[i].name = name
		T[i].sid = sid
		T[i].gpa = gpa
		fmt.Scanln(&name, &sid, &gpa)
		i++
		*n++
	}
}

func sort(M *tabMhs, n int) {
	var pass, idx, i int
	var temp student
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if M[idx].gpa > M[i].gpa {
				idx = i
			}
			i++
		}

		// T[idx], T[pass-1] = T[pass-1], T[idx]
		temp = M[pass-1]
		M[pass-1] = M[idx]
		M[idx] = temp
		pass = pass + 1
	}
}

func cetakArray(M tabMhs, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s %s %g\n", i+1, M[i].sid, M[i].name, M[i].gpa)
	}
}

func main() {
	var mahasiswa tabMhs
	var n int

	input(&mahasiswa, &n)
	sort(&mahasiswa, n)
	cetakArray(mahasiswa, n)
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
