package main

import "fmt"

const NMAX = 1001

type Pemain struct {
	nama   string
	gol    int
	assist int
}

type arrPemain [NMAX]Pemain

func InsertionSort(T *arrPemain, n int) {
	for i := 1; i < n; i++ {
		key := T[i]
		j := i - 1
		for j >= 0 && (T[j].gol < key.gol || (T[j].gol == key.gol && T[j].assist < key.assist)) {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = key
	}
}

func main() {
	var A arrPemain
	var n int

	fmt.Println("Masukan Data Input :")
	fmt.Scan(&n)

	var nama1, nama2 string
	for i := 0; i < n; i++ {
		fmt.Scan(&nama1, &nama2, &A[i].gol, &A[i].assist)
		A[i].nama = nama1 + " " + nama2
	}

	InsertionSort(&A, n)

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %d\n", A[i].nama, A[i].gol, A[i].assist)
	}
}