package main

import "fmt"

func isiArray(t *tabel, n *int) {
	var temp string

	fmt.Scan(&temp)
	for temp != "." {
		t[*n] = temp
		fmt.Scan(&temp)
		*n++
	}
}

//func cetakArray(t tabel, n int){
//	var i int
//
//	for i < &n{
//		fmt.Print(t[i])
//		i++
//	}
//}

func balikanArray(t *tabel, n int) {
	for i := 0; i <= (n-1)/2; i++ {
		temp := t[i]
		t[i] = t[(n-1)-i]
		t[(n-1)-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	titip := t
	balikanArray(&titip, n)
	return titip == t
}

const NMAX = 127

type tabel = [NMAX]string

func main() {

	var t tabel
	var n int
	// var temp string

	n = 0

	isiArray(&t, &n)
	//	cetakArray(t,n)
	//	balikanArray(&t,&n,&temp)
	fmt.Println(palindrom(t, n))
}
