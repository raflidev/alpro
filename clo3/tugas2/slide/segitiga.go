package main

import "fmt"

func main() {
	var s1, s2, s3 int
	var segitiga bool
	fmt.Scanln(&s1, &s2, &s3)

	segitiga = cekSegitiga(s1, s2, s3)

	if segitiga {
		fmt.Println("Segitiga")
	} else {
		fmt.Println("Bukan segitiga")

	}

}

func cekSegitiga(s1, s2, s3 int) bool {
	if s1+s2 <= s3 || s1+s3 <= s2 || s2+s3 <= s1 {
		return false
	} else {
		return true
	}
}
