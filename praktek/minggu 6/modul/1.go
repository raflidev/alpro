package main

import "fmt"

func main() {
	var n, r, result int
	fmt.Scanln(&n, &r)
	findFactorial(n, &result)
	fmt.Println(Permutation(n, r), Combination(n, r))
}
func findFactorial(n int, result *int) {
	var j int
	j = 1
	for i := 1; i <= n; i++ {
		j = j * i
	}
	*result = j
}

func Permutation(n, r int) int {
	var res1, res2 int
	findFactorial(n, &res1)
	findFactorial(n-r, &res2)
	return res1 / res2
}

func Combination(n, r int) int {
	var res1, res2, res3 int
	findFactorial(n, &res1)
	findFactorial(r, &res2)
	findFactorial(n-r, &res3)
	return res1 / (res2 * res3)
}
