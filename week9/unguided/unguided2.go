package main

import (
	"fmt"
	"math"
)

const MAX = 100

func tampilArray(arr [MAX]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func rataRata(arr [MAX]int, n int) float64 {
	var total int

	for i := 0; i < n; i++ {
		total += arr[i]
	}

	return float64(total) / float64(n)
}

func standarDeviasi(arr [MAX]int, n int) float64 {
	mean := rataRata(arr, n)
	var jumlahKuadrat float64

	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - mean
		jumlahKuadrat += selisih * selisih
	}

	return math.Sqrt(jumlahKuadrat / float64(n))
}

func frekuensi(arr [MAX]int, n int, bilangan int) int {
	var jumlah int

	for i := 0; i < n; i++ {
		if arr[i] == bilangan {
			jumlah++
		}
	}

	return jumlah
}

func hapusElemen(arr *[MAX]int, n *int, indeks int) {
	for i := indeks; i < *n-1; i++ {
		arr[i] = arr[i+1]
	}

	*n--
}

func main() {
	var arr [MAX]int
	var n int
	var x int
	var indeksHapus int
	var bilanganCari int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&x)
	fmt.Scan(&indeksHapus)
	fmt.Scan(&bilanganCari)

	fmt.Print("Isi array: ")
	tampilArray(arr, n)

	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	hapusElemen(&arr, &n, indeksHapus)

	fmt.Print("Isi array setelah penghapusan: ")
	tampilArray(arr, n)

	fmt.Printf("Rata-rata: %.2f\n", rataRata(arr, n))
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi(arr, n))

	fmt.Println("Frekuensi", bilanganCari, ":", frekuensi(arr, n, bilanganCari))
}