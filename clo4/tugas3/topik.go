package main

import "fmt"

const nmax = 100

type tag struct {
	topik  string
	banyak int
}

type tabTopic = [nmax]string

type tabTag = [nmax]tag

func isiTopik(T *tabTopic, n *int) {
	var data string
	fmt.Scan(&data)
	i := 0
	for i < nmax && data != "#" {
		T[i] = data
		fmt.Scan(&data)
		*n++
		i++
	}
}
func cariBanyakTopik(T tabTopic, n int, topik string, banyak *int) {
	*banyak = 0
	for i := 0; i < n; i++ {
		if T[i] == topik {
			*banyak++
		}
	}
}
func isiTag(T tabTopic, n int, tab *tabTag) {
	j, i := 0, 0
	for j < n {
		if T[j] != "" {
			banyak := 0
			topik := T[j]
			cariBanyakTopik(T, n, topik, &banyak)
			tab[j].topik = topik
			tab[j].banyak = banyak
			for i < n {
				if T[i] == topik {
					T[i] = ""
				}
				i++
			}
		}
		j++
		i = 0
	}
}

func cariTrendingTopik(T tabTag, n int) {
	max, idx := 0, 0
	for i := 0; i < n; i++ {
		if T[i].topik != "" {
			if max < T[i].banyak {
				max = T[i].banyak
				idx = i
			}
		}
	}

	fmt.Println(T[idx].topik)
}

func main() {
	var topik tabTopic
	var tabTag tabTag
	var n int
	isiTopik(&topik, &n)
	isiTag(topik, n, &tabTag)
	cariTrendingTopik(tabTag, n)
	// fmt.Println(tabTag)
}

// libur lebaran libur tugas libur lebaran libur mudik libur libur rendang libur puasa uas libur lebaran tugas libur libur 2minggu vaksin #
