package main

import "fmt"

func posisi(tab [107]string, n int, x string) int {
	var index int
	i := 0
	found := false
	for i < n && !found {
		if tab[i] == x {
			index = i
			found = true
		}
		i++
	}
	return index
}

func hapus(tab *[107]string, n *int, x string) {
	idx := posisi(*tab, *n, x)
	for i := idx; i < *n; i++ {
		tab[i] = tab[i+1]
	}
	*n--
}

func updateKelulusan(mhs *[107]string, wisuda [107]string, n *int, m int) {
	for i := 0; i < m; i++ {
		for j := 0; j < *n; j++ {
			if mhs[j] == wisuda[i] {
				hapus(mhs, &*n, mhs[i])
			}
		}
	}
}

func main() {
	var i, n, m int
	var daftar, wisuda [107]string
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&daftar[i])
	}
	fmt.Scan(&m)
	for i = 0; i < m; i++ {
		fmt.Scan(&wisuda[i])
	}

	updateKelulusan(&daftar, wisuda, &n, m)

	fmt.Print(n, " ")
	for i = 0; i < n; i++ {
		fmt.Print(daftar[i], " ")
	}
	fmt.Print("\n")
}

// 9 130440 130939 130210 130879 130793 130943 130531 130823 130879
// 6 130939 130531 130440 130141 130943 130879
