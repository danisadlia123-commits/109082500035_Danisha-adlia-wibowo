package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	var totalSemua float64

	for i := 0; i < jumlahWadah; i++ {
		var totalWadah float64

		for j := i * y; j < (i+1)*y && j < x; j++ {
			totalWadah += berat[j]
		}

		totalSemua += totalWadah
		fmt.Printf("%g ", totalWadah)
	}

	rataRata := totalSemua / float64(jumlahWadah)

	fmt.Println()
	fmt.Printf("%g\n", rataRata)
}