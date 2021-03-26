package main

import "fmt"

func main() {
	var i, x, f1, maxi, mini int
	fmt.Scan(&x)
	// looping inputan dan mencari value
	for i = 1; i <= x; i++ {
		fmt.Scan(&f1)

		min(f1, &mini)
		max(f1, &maxi)
	}
	if mini == 0 {
		fmt.Println(maxi, maxi)
	} else {
		fmt.Println(maxi, mini)
	}
}

func min(f1 int, f2 *int) {
	if f1 <= *f2 {
		*f2 = f1
	}
}

func max(f1 int, f2 *int) {
	if f1 >= *f2 {
		*f2 = f1
	}
}
