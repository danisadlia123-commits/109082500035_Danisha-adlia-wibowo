package main

import "fmt"

const nMax int = 1000000

type arrInt [nMax]int

func insertionSort(T *arrInt, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func main() {
	var data arrInt
	var n int
	var input int

	n = 0
	fmt.Scan(&input)
	for input >= 0 {
		data[n] = input
		n++
		fmt.Scan(&input)
	}

	insertionSort(&data, n)

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Print(data[i])
		} else {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	if n <= 1 {
		fmt.Println("Data berjarak 0")
	} else {
		jarak := data[1] - data[0]
		tetap := true
		for i := 2; i < n; i++ {
			if data[i]-data[i-1] != jarak {
				tetap = false
				break
			}
		}
		if tetap {
			fmt.Println("Data berjarak", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}