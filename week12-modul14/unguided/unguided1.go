package main

import "fmt"

const nMax int = 1000000

type arrInt [nMax]int

func selectionSort(T *arrInt, n int) {
	var idxMin, i, j int
	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if T[j] < T[idxMin] {
				idxMin = j
			}
		}
		if idxMin != i {
			T[i], T[idxMin] = T[idxMin], T[i]
		}
	}
}

func main() {
	var n, m int
	var data arrInt

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&data[j])
		}

		selectionSort(&data, m)

		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Print(data[j])
			} else {
				fmt.Print(data[j], " ")
			}
		}
		fmt.Println()
	}
}