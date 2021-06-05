package main

import "fmt"

const nmax = 60

type partai struct {
	nama, suara int
}

type tabPartai = [nmax]partai

func main() {
	var p tabPartai
	var input, n, found int
	n = 0

	fmt.Scan(&input)
	for input != -1 {
		found = posisi(p, n, input)
		if found != -1 {
			p[found].suara++
		} else {
			p[n].nama = input
			p[n].suara++
			n++
		}
		fmt.Scan(&input)
	}

	for i := 0; i < n; i++ {
		j := i + 1
		for j > 0 && p[j-1].suara < p[j].suara {
			temp := p[j]
			p[j] = p[j-1]
			p[j-1] = temp
			j--
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
}

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

// 5 1 1 1 1 1 1 1 3 3 3 3 3 2 2 5 5 5 5 5 4 3 2 2 2 2 -1
