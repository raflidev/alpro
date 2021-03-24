package main

import "fmt"

func main() {

	var key, tb1, tb2, tb3 int

	fmt.Scanln(&key)
	fmt.Scanln(&tb1)
	fmt.Scanln(&tb2)
	fmt.Scanln(&tb3)

	k4 := key & 10
	k3 := key / 10 & 10
	k2 := key / 100 & 10
	k1 := key / 1000

	q3 := key & 100
	q2 := key / 10 & 100
	q1 := key / 100

	// pegecheckan
	fmt.Println(tb1 == key)
	fmt.Println(k1 == tb1 || k2 == tb1 || k3 == tb1 || k4 == tb1)
	fmt.Println(q1 == tb2 || q2 == tb2 || q3 == tb2)
}
