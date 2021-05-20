package main

import "fmt"

const nmax = 1000

type waktu struct {
	menit, detik int
}

type lagu struct {
	judul, penyanyi string
	waktu
}

type playlist = [nmax]lagu

func cariLagu(T *playlist, n int) {
	i, j := 0, 0
	for i < n {
		for j < n {
			if T[i].judul == T[j].judul && T[i].penyanyi == T[j].penyanyi && j != i {
				T[j].judul = ""
				T[j].penyanyi = ""
				T[j].menit = 0
				T[j].detik = 0
			}
			j++
		}
		i++
		j = 0
	}
}

func buatPlaylist(T *playlist, n *int) {
	var judul, penyanyi string
	var menit, detik int

	fmt.Scanln(&judul, &penyanyi, &menit, &detik)
	i := 0
	for (judul != "#" && penyanyi != "#") && i < nmax {
		T[i].judul = judul
		T[i].penyanyi = penyanyi
		T[i].menit = menit
		T[i].detik = detik
		fmt.Scanln(&judul, &penyanyi, &menit, &detik)
		*n = *n + 1
		i++
	}
}

func laguTerlama(T *playlist, n int) {
	var idx, max int
	idx = 0
	max = (T[0].menit * 60) + T[0].detik
	for i := 0; i < n; i++ {
		if (T[i].menit*60)+T[i].detik > max {
			max = (T[i].menit * 60) + T[i].detik
			idx = i
		}
	}
	T[idx].judul = "*" + T[idx].judul
}

func cetakPlaylist(T *playlist, n int) {
	for i := 0; i < n; i++ {
		if T[i].judul != "" {
			if string([]rune(T[i].judul)[0]) == "*" {
				fmt.Println(T[i].judul, T[i].menit, "menit", T[i].detik, "detik")
			} else {
				fmt.Println(T[i].judul)
			}
		}
	}
}

func main() {
	var list playlist
	var n int
	n = 0
	buatPlaylist(&list, &n)
	cariLagu(&list, n)
	laguTerlama(&list, n)
	cetakPlaylist(&list, n)
	// fmt.Println(list)
}

// snowman sia 2 53
// bertaubat nadinamizah 5 21
// anyone justinbieber 3 16
// cuek rizkyfebian 4 20
// anyone justinbieber 4 24
// sofia clairo 3 8
// snowman sia 2 42
// # #

// somebodytoldme thekiller 3 21
// somebodytoldme thekiller 3 22
// somebodytoldme thekiller 3 23
// somebodytoldme thekiller 3 24
// # #
