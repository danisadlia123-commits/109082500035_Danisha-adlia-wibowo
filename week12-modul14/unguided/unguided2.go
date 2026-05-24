package main

import "fmt"

const nMax int = 1000000

type arrInt [nMax]int

func selectionSortAsc(T *arrInt, n int) {
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

func selectionSortDesc(T *arrInt, n int) {
	var idxMax, i, j int
	for i = 0; i < n-1; i++ {
		idxMax = i
		for j = i + 1; j < n; j++ {
			if T[j] > T[idxMax] {
				idxMax = j
			}
		}
		if idxMax != i {
			T[i], T[idxMax] = T[idxMax], T[i]
		}
	}
}

func main() {
	var n, m int
	var data arrInt
	var ganjil, genap arrInt
	var nGanjil, nGenap int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		nGanjil = 0
		nGenap = 0

		for j := 0; j < m; j++ {
			fmt.Scan(&data[j])
			if data[j]%2 != 0 {
				ganjil[nGanjil] = data[j]
				nGanjil++
			} else {
				genap[nGenap] = data[j]
				nGenap++
			}
		}

		selectionSortAsc(&ganjil, nGanjil)
		selectionSortDesc(&genap, nGenap)

		for j := 0; j < nGanjil; j++ {
			fmt.Print(ganjil[j], " ")
		}
		for j := 0; j < nGenap; j++ {
			if j == nGenap-1 {
				fmt.Print(genap[j])
			} else {
				fmt.Print(genap[j], " ")
			}
		}
		fmt.Println()
	}
}