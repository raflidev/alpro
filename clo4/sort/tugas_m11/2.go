package main

import "fmt"

const MAX = 30

type intArray struct {
	tabInt [MAX]int
	N      int
}

var array1, array2 intArray

func inputArray(T *intArray) {
	var data, i int
	i = 0
	fmt.Scan(&data)
	for i < MAX && data != 0 {
		T.tabInt[i] = data
		T.N++
		fmt.Scan(&data)
		i++
	}
}

func sortArray(T *intArray) {
	var pass, idx, i, temp int
	pass = 1
	for pass <= T.N-1 {
		idx = pass - 1
		i = pass
		for i < T.N {
			if T.tabInt[idx] > T.tabInt[i] {
				idx = i
			}
			i++
		}
		temp = T.tabInt[pass-1]
		T.tabInt[pass-1] = T.tabInt[idx]
		T.tabInt[idx] = temp
		pass++
	}
}

func isSimiliar(T1, T2 intArray) bool {
	fmt.Println(T1, T2)
	return T1.tabInt == T2.tabInt
}

func main() {
	inputArray(&array1)
	inputArray(&array2)
	sortArray(&array1)
	sortArray(&array2)
	fmt.Println(isSimiliar(array1, array2))
}

// 1 2 4 3 5 0
// 1 2 3 5 4 0
