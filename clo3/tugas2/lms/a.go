package main

import "fmt"

func main() {
	var a int
	a = 2
	tambah(&a)
	fmt.Println(a)

}
func tambah(a *int) {
	*a += 5
}
