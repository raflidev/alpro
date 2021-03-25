package main

import "fmt"

func main() {
	var bil1, bil2 int
	var x, y, z int
	fmt.Scanln(&bil1, &bil2)
	x = P(bil1)
	y = P(bil2)
	z = P(bil1 - bil2)

	fmt.Println(x, y, x/z)

}
func P(x int) int {
	var j int
	j = 1
	for i := 1; i <= x; i++ {
		j = j * i
	}
	return j
}
