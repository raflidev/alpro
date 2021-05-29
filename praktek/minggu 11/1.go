package main

import "fmt"

const NMAX = 1000

type lagu struct {
	judul, penyanyi string
	durasi          waktu
}

type waktu struct {
	menit, detik int
}

type playlist [NMAX]lagu

func cariLagu(T playlist, n int, judul, penyanyi string) bool {
	for i := 0; i < n; i++ {
		if T[i].judul == judul && T[i].penyanyi == penyanyi {
			return true
		}
	}
	return false
}

func buatPlaylist(T *playlist, n *int) {
	var judul, penyanyi string
	var menit, detik int
	var duplikat bool
	fmt.Scanln(&judul, &penyanyi, &menit, &detik)
	i := 0
	for (judul != "#" && penyanyi != "#") && i < NMAX {
		T[i].judul = judul
		T[i].penyanyi = penyanyi
		T[i].durasi.menit = menit
		T[i].durasi.detik = detik
		duplikat = cariLagu(*T, i, T[i].judul, T[i].penyanyi)
		if duplikat {
			T[i].judul = ""
			T[i].penyanyi = ""
			T[i].durasi.menit = 0
			T[i].durasi.detik = 0
		}
		fmt.Scanln(&judul, &penyanyi, &menit, &detik)
		*n = *n + 1
		i++
	}
}

func terlama(T *playlist, n int) {
	var idx, max int
	idx = 0
	max = 0
	for i := 0; i < n; i++ {
		if (T[i].durasi.menit*60)+T[i].durasi.detik > max {
			max = (T[i].durasi.menit * 60) + T[i].durasi.detik
			idx = i
		}
	}
	T[idx].judul = "*" + T[idx].judul
}

func cetakPlaylist(T playlist, n int) {
	for i := 0; i < n; i++ {
		if T[i].judul != "" {
			if string([]rune(T[i].judul)[0]) == "*" {
				fmt.Println(T[i].judul, T[i].durasi.menit, "menit", T[i].durasi.detik, "detik")
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
	terlama(&list, n)
	cetakPlaylist(list, n)

}
