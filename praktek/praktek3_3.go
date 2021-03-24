package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5 int
	var s1, s2, s3, s4, s5 int
	var straight, flush bool
	fmt.Scanf("%d%c %d%c %d%c %d%c %d%c", &a1, &s1, &a2, &s2, &a3, &s3, &a4, &s4, &a5, &s5)
	straight = (a1+1 == a2 && a2+1 == a3 && a3+1 == a4 && a4+1 == a5 && a1 != 1) || (a1 == 10 && a2 == 11 && a3 == 12 && a4 == 13 && a5 == 1)
	flush = s1 == s2 && s2 == s3 && s3 == s4 && s4 == s5

	fmt.Println(straight)
	fmt.Println(flush)
}
