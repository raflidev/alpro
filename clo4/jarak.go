package main

import (
	"fmt"
	"math"
)

type titik struct {
	t1, t2 float64
}

func main() {
	var P1, P2 titik

	fmt.Scanln(&P1.t1, &P1.t2, &P2.t1, &P2.t2)
	fmt.Println(jarak(P1, P2))
}

func jarak(P1, P2 titik) float64 {
	return akar(math.Pow((P1.t1-P2.t1), 2) + math.Pow((P1.t2-P2.t2), 2))
}

func akar(x float64) float64 {
	return math.Sqrt(x)
}
