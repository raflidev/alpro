// vpl

package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	for i := 0; i <= b-a; i++ {
		x := a + i
		fmt.Printf("Simbol ASCII dari %d adalah %c \n", x, x)
	}
}
